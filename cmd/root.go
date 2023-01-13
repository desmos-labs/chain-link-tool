package cmd

import (
	chainlinkgetter "github.com/desmos-labs/desmos/v4/app/desmos/cmd/chainlink/getter"
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command of the application
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-link-tool",
		Short: "Chain link tool is a tool that generates a proof json for desmos chain link",
	}
	cmd.AddCommand(
		NewGenerateCommand(chainlinkgetter.NewChainLinkReferencePrompt()),
	)
	return cmd
}
