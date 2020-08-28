package app

import (
	"github.com/spf13/cobra"
)

var whoamiCommand = &cobra.Command{
	Use:     "whoami",
	Aliases: []string{"w"},
	Short:   "print about yourself",
	RunE:    whoamiCommandRunE,
}

func whoamiCommandRunE(cmd *cobra.Command, args []string) error {
	cmd.Printf("user\t%s\t%s\n", userName, userID)
	cmd.Printf("bot\t%s\t%s\n", botName, botID)

	return nil
}

func init() {
	RootCommand.AddCommand(whoamiCommand)
	whoamiCommand.PersistentFlags().BoolP("identifier", "i", false, "print user id instead of name")
}
