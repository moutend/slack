package app

import (
	"fmt"

	"github.com/moutend/slack/internal/utility"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var messageCommand = &cobra.Command{
	Use:     "message",
	Aliases: []string{"m", "messages"},
	Short:   "print messages",
	RunE:    messageCommandRunE,
}

func messageCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	forceFetch, _ := cmd.Flags().GetBool("debug")
	target := args[0]

	users, err := api.GetAllUsersContext(cmd.Context())

	if err != nil {
		return err
	}

	userIdToName := make(map[string]string)

	for _, user := range users {
		userIdToName[user.ID] = user.Name

		if "@"+user.Name == target {
			target = user.ID
		}
	}

	channels, err := api.GetAllChannelsContext(cmd.Context())

	if err != nil {
		return err
	}

	var found slack.Channel

	for _, c := range channels {
		if c.Name == target || c.NameNormalized == target || c.User == target {
			found = c

			break
		}
	}
	if found.ID == "" {
		return fmt.Errorf("channel or user %s not found", target)
	}

	messages, err := api.GetAllConversationsContext(cmd.Context(), found.ID, forceFetch)

	if err != nil {
		return err
	}
	if err := utility.SortMessagesByTimestamp(messages); err != nil {
		return err
	}

	for _, m := range messages {
		u, ok := userIdToName[m.User]

		if !ok {
			u = "undefined"
		}

		cmd.Println(
			"@"+u,
			utility.FormatMessageText(m.Text, userIdToName),
			utility.Ago(m.Timestamp),
		)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(messageCommand)
}
