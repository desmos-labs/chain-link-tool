package cmd

import (
	"fmt"
	"os"

	"github.com/desmos-labs/desmos/v4/app"

	"github.com/desmos-labs/chain-link-tool/cmd/provider"
	"github.com/desmos-labs/chain-link-tool/cmd/types"
	"github.com/spf13/cobra"
)

func NewGenerateCommand(getter types.ChainLinkReferenceGetter) *cobra.Command {
	return &cobra.Command{
		Use:   "create-chain-link-json [provider]",
		Short: "Start an interactive prompt to create a new chain link JSON object",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get Desmos link owner
			owner, err := getter.GetOwner()
			if err != nil {
				return err
			}

			// Generate chain link JSON by the given provider
			provider, err := provider.DefaultChainLinkJSONBuilderProvider(owner, args[0])
			if err != nil {
				return err
			}
			chainLinkJSON, err := provider.BuildChainLinkJSON()
			if err != nil {
				return err
			}

			// Marshal the chain link JSON
			bz, err := app.MakeTestEncodingConfig().Marshaler.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			// Write the chain link JSON to a file
			filename, err := getter.GetFilename()
			if err != nil {
				return err
			}
			err = os.WriteFile(filename, bz, 0600)
			if err != nil {
				return err
			}

			cmd.Println(fmt.Sprintf("Chain link JSON file stored at %s", filename))
			return nil
		},
	}
}
