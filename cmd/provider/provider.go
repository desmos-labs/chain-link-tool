package provider

import (
	"fmt"

	"github.com/desmos-labs/chain-link-tool/cmd/nomic"
	"github.com/desmos-labs/chain-link-tool/cmd/types"
)

// DefaultChainLinkJSONBuilderProvider returns the default ChainLinkJSONBuilder provider implementation
func DefaultChainLinkJSONBuilderProvider(owner string, name string) (types.ChainLinkJSONBuilder, error) {
	switch name {
	case nomic.ChainName:
		return nomic.NewDefaultNomicBuilder(owner), nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", name)
	}
}
