package app

import (
	"github.com/moutend/slack/internal/version"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "print version",
	RunE:    versionCommandRunE,
}

func versionCommandRunE(cmd *cobra.Command, args []string) error {
	cmd.Println(version.String())

	return nil
}

func init() {
	RootCommand.AddCommand(versionCommand)
}
