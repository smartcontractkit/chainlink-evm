package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/nonce_manager"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// Config represents the overall structure of the YAML.
type Config struct {
	ChainSelector string   `yaml:"chain_selector"`
	ChainID       string   `yaml:"chain_id"`
	Type          string   `yaml:"type"`
	Names         Names    `yaml:"names"`
	Settings      Settings `yaml:"settings"`
	Nodes         []Node   `yaml:"nodes"`
}

// Names contains the different naming conventions for the blockchain.
type Names struct {
	Blockchain string `yaml:"blockchain"`
	Network    string `yaml:"network"`
	Shortname  string `yaml:"shortname"`
	Fullname   string `yaml:"fullname"`
	Canonical  string `yaml:"canonical"`
	RMN        string `yaml:"rmn"`
}

// Settings holds various configuration settings.
type Settings struct {
	ClNodes ClNodes `yaml:"clNodes"`
	Proxy   Proxy   `yaml:"proxy"`
}

// ClNodes contains Chainlink node-specific settings.
type ClNodes struct {
	UseRpcProxy bool `yaml:"useRpcProxy"`
}

// Proxy contains proxy-specific settings.
type Proxy struct {
	HealthCheckType            string `yaml:"healthCheckType"`
	HealthCheckInterval        int    `yaml:"healthCheckInterval"`
	HealthCheckBlockDifference int    `yaml:"healthCheckBlockDifference"`
}

// Node represents a single node configuration.
type Node struct {
	Provider  string   `yaml:"provider"`
	HTTPURL   string   `yaml:"httpUrl"`
	WSURL     string   `yaml:"wsUrl"`
	Comment   string   `yaml:"comment"`
	Order     int      `yaml:"order"`
	Dedicated string   `yaml:"dedicated"` // Using string because it's "false" or "true" in YAML
	Tags      []string `yaml:"tags"`
}

func readAllRPCConfigs(folderPath string) ([]Config, error) {
	// Read all entries in the directory
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}
	var result []Config
	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}

		fileName := file.Name()
		fileExtension := strings.ToLower(filepath.Ext(fileName))

		if fileExtension != ".yaml" && fileExtension != ".yml" {
			continue
		}

		filePath := filepath.Join(folderPath, fileName)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		var rpcCfg Config
		err = yaml.Unmarshal(content, &rpcCfg)
		if err != nil {
			log.Fatalf("Error unmarshaling YAML: %v", err)
		}

		result = append(result, rpcCfg)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Names.Canonical < result[j].Names.Canonical
	})
	return result, nil
}

type Contract struct {
	Type string
}

func readContracts(path string, result map[string]string) error {
	contractsByChainSelector := map[string]map[string]Contract{}
	jsonFile, err := ioutil.ReadFile(path) // Assuming your YAML is in a file named config.yaml
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	err = json.Unmarshal(jsonFile, &contractsByChainSelector)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	for rawChainID, contracts := range contractsByChainSelector {
		for contractAddr, contract := range contracts {
			if contract.Type == "NonceManager" || contract.Type == "KeystoneForwarder" {
				result[rawChainID] = contractAddr
				break
			}
		}
	}

	return nil
}

type Result struct {
	ChainName string
	RPCName   string
	InMinutes float64
	Error     error
}

func TestCall(t *testing.T) {
	rpcConfigToChainSelector, err := readAllRPCConfigs("/Users/dh/infra-k8s/rpcs/networks.d/production/")
	require.NoError(t, err)
	lggr := logger.Test(t)
	contracts := map[string]string{}
	require.NoError(t, readContracts("/Users/dh/go/src/github.com/smartcontractkit/chainlink-deployments/domains/ccip/mainnet/addresses.json", contracts))
	require.NoError(t, readContracts("/Users/dh/go/src/github.com/smartcontractkit/chainlink-deployments/domains/ccip/staging/addresses.json", contracts))
	require.NoError(t, readContracts("/Users/dh/go/src/github.com/smartcontractkit/chainlink-deployments/domains/keystone/mainnet/addresses.json", contracts))

	var results []Result
	for _, rpcConfig := range rpcConfigToChainSelector {
		contract, ok := contracts[rpcConfig.ChainSelector]
		if !ok {
			lggr.Infof("No contract found. Skipping %s.", rpcConfig.Names.Canonical)
			continue
		}

		subR := ensureSufficientLookBackForChain(t, rpcConfig, contract)
		results = append(results, subR...)
	}

	for _, result := range results {
		println(fmt.Sprintf("%s\t%s\t%.2f\t%v", result.ChainName, result.RPCName, result.InMinutes, result.Error))
	}

}

func ensureSufficientLookBackForChain(t *testing.T, rpcCfg Config, contractAddr string) []Result {
	lggr := logger.Test(t)
	name := rpcCfg.Names.Fullname
	var results []Result
	for _, node := range rpcCfg.Nodes {
		lookBack, err := findLookBackDuration(node.HTTPURL, contractAddr)
		if err != nil {
			lggr.Errorf("Error finding look back duration for %s using %s: %v", name, node.HTTPURL, err)
		}

		if lookBack < time.Hour {
			lggr.Warnf("lookback for %s %s is not old enought %s", name, node.HTTPURL, lookBack)
		} else {
			lggr.Infof("lookback for %s %s is older than %s", name, node.HTTPURL, lookBack)
		}

		results = append(results, Result{
			ChainName: name,
			RPCName:   node.Provider,
			InMinutes: float64(lookBack) / float64(time.Minute),
			Error:     err,
		})
	}

	return results
}

func headByNumber(client *ethclient.Client, number *big.Int) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return client.HeaderByNumber(ctx, number)

}

func findLookBackDuration(rpcURL, contract string) (dur time.Duration, err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err = recovered.(error)
		}
	}()
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return 0, fmt.Errorf("failed to dial client: %w", err)
	}
	client.Close()
	latestBlock, err := headByNumber(client, nil)
	if err != nil {
		return 0, fmt.Errorf("error getting latest block number: %v", err)
	}

	addr := common.HexToAddress(contract)
	registry, err := nonce_manager.NewNonceManager(addr, client)
	if err != nil {
		return 0, fmt.Errorf("error getting capabilities registry: %v", err)
	}

	const maxLookBack = 2000
	startTime := time.Unix(int64(latestBlock.Time), 0)
	var maxDuration time.Duration
	_, found := sort.Find(maxLookBack, func(offset int) int {
		blockNumToTry := latestBlock.Number.Int64() - (maxLookBack - int64(offset))

		block, err := headByNumber(client, big.NewInt(blockNumToTry))
		if err != nil {
			panic(fmt.Errorf("error getting block number: %v", err))
		}

		_, err = registry.Owner(&bind.CallOpts{
			BlockNumber: big.NewInt(blockNumToTry),
			Context:     nil,
		})
		if err != nil {
			if strings.Contains(err.Error(), "missing trie node") {
				return 1 // move offset to more resent block
			}
			panic(fmt.Errorf("failed to call owner due to unknow error: %w", err))
		}

		maxDuration = startTime.Sub(time.Unix(int64(block.Time), 0))

		return 0
	})

	if !found {
		return 0, fmt.Errorf("failed to find look back for %s", contract)
	}

	return maxDuration, nil
}
