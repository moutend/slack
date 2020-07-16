package app

import (
	"sort"

	"github.com/spf13/cobra"
)

var channelCommand = &cobra.Command{
	Use:     "channel",
	Aliases: []string{"c", "channels"},
	Short:   "print available channels",
	RunE:    channelCommandRunE,
}

func channelCommandRunE(cmd *cobra.Command, args []string) error {
	users, err := api.GetAllUsersContext(cmd.Context())

	if err != nil {
		return err
	}

	usersMap := make(map[string]struct{})
	userIdToName := make(map[string]string)

	for _, user := range users {
		userIdToName[user.Name] = user.ID
	}

	channels, err := api.GetAllChannelsContext(cmd.Context())

	if err != nil {
		return err
	}

	names := []string{}

	for _, channel := range channels {
		usersMap[channel.User] = struct{}{}

		if channel.Name != "" {
			names = append(names, channel.Name)
		}
	}

	userNames := []string{}

	for k, _ := range usersMap {
		if k == "" {
			continue
		}

		userName, _ := userIdToName[k]

		if userName == "" {
			continue
		}

		userNames = append(userNames, userName)
	}

	sort.Strings(userNames)
	sort.Strings(names)

	for _, userName := range userNames {
		cmd.Printf("@%s\n", userName)
	}
	for _, name := range names {
		cmd.Println(name)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(channelCommand)
}
