package types

import (
	"github.com/desmos-labs/desmos/v4/x/profiles/client/utils"
)

type ChainLinkReferenceGetter interface {
	GetFilename() (string, error)

	GetOwner() (string, error)
}

type ChainLinkJSONBuilder interface {
	BuildChainLinkJSON() (utils.ChainLinkJSON, error)
}
