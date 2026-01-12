package gobindings

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// VersionHash is the hash used to detect changes in the underlying contract
func VersionHash(abiPath string, binPath string) (hash string) {
	abi, err := os.ReadFile(abiPath)
	if err != nil {
		Exit("Could not read abi path to create version hash", err)
	}
	bin := []byte("")
	if binPath != "-" {
		bin, err = os.ReadFile(binPath)
		if err != nil {
			Exit("Could not read bin path to create version hash", err)
		}
	}
	hashMsg := string(abi) + string(bin) + "\n"
	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashMsg)))
}

func Exit(msg string, err error) {
	if err != nil {
		fmt.Println(msg+":", err)
	} else {
		fmt.Println(msg)
	}
	os.Exit(1)
}
