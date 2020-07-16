package app

import (
	"sort"

	"github.com/spf13/cobra"
)

var userCommand = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u", "users"},
	Short:   "print users",
	RunE:    userCommandRunE,
}

func userCommandRunE(cmd *cobra.Command, args []string) error {
	users, err := api.GetAllUsersContext(cmd.Context())

	if err != nil {
		return err
	}

	names := []string{}

	for _, user := range users {
		if user.Name != "" {
			names = append(names, user.Name)
		}
	}

	sort.Strings(names)

	for _, name := range names {
		cmd.Printf("@%s\n", name)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(userCommand)
}
