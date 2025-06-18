package bindings

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-evm/pkg/bindings/abigen"
)

// Generate returns the Go source for your contract bindings.
func Generate(opts Options) (string, error) {
	var (
		code string
		err  error
	)

	code, err = abigen.BindV2(
		opts.Types,
		opts.ABIs,
		opts.Bins,
		opts.Package,
		opts.Libs,
		opts.Aliases,
		tpl,
	)

	if err != nil {
		return "", fmt.Errorf("bindgen: %w", err)
	}
	return code, nil
}
