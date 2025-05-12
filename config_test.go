package chainlink_evm

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

var (
	//go:embed CONFIG.md
	configMD string
)

//go:generate go run ./pkg/scmd/config-docs
func TestConfigDocs(t *testing.T) {
	cfg, err := toml.GenerateDocs()
	assert.NoError(t, err, "invalid config docs")
	assert.Equal(t, configMD, cfg, "CONFIG.md is out of date. Run 'go generate .' to regenerate.")
}
