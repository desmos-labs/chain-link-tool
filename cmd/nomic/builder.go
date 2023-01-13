package nomic

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/desmos-labs/chain-link-tool/cmd/types"
)

const (
	ChainName   = "nomic"
	chainPrefix = "nomic"
)

var _ types.ChainLinkJSONBuilder = &NomicBuilder{}

type NomicBuilder struct {
	owner  string
	getter NomicGetter
}

func NewNomicBuilder(owner string, getter NomicGetter) *NomicBuilder {
	return &NomicBuilder{
		owner:  owner,
		getter: getter,
	}
}

func NewDefaultNomicBuilder(owner string) *NomicBuilder {
	return &NomicBuilder{
		owner:  owner,
		getter: &nomicGetter{},
	}
}

func (b *NomicBuilder) BuildChainLinkJSON() (utils.ChainLinkJSON, error) {
	keyPath, err := b.getter.GetPrivateKeyPath()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	privateKey, err := parseNomicPrivateKeyFromFile(keyPath)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	value := []byte(b.owner)
	sig, err := privateKey.Sign(value)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	addr, err := sdk.Bech32ifyAddressBytes(chainPrefix, privateKey.PubKey().Address())
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	return utils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chainPrefix),
		profilestypes.NewProof(
			privateKey.PubKey(),
			profilestypes.NewSingleSignature(profilestypes.SIGNATURE_VALUE_TYPE_RAW, sig),
			hex.EncodeToString(value),
		),
		profilestypes.NewChainConfig(ChainName),
	), nil
}

func parseNomicPrivateKeyFromFile(file string) (cryptotypes.PrivKey, error) {
	keyBz, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(keyBz) != secp256k1.PrivKeySize {
		return nil, fmt.Errorf("invalid private key length")
	}

	return &secp256k1.PrivKey{
		Key: keyBz,
	}, nil
}
