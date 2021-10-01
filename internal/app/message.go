package app

import (
	"context"
	"strings"

	"github.com/moutend/slack/internal/models"
	"github.com/moutend/slack/internal/utility"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var messageCommand = &cobra.Command{
	Use:     "message",
	Aliases: []string{"m"},
	Short:   "print message",
	RunE:    messageCommandRunE,
}

func messageCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	var userNameReplacer *strings.Replacer
	var messages []*models.Message

	err := dbm.Transaction(cmd.Context(), func(ctx context.Context, tx boil.ContextTransactor) error {
		if yes, _ := cmd.Flags().GetBool("offline"); yes {
			goto LOAD_USER_CACHE
		}
		if err := client.FetchUsers(ctx, tx); err != nil {
			return err
		}
		if err := client.FetchChannels(ctx, tx); err != nil {
			return err
		}

	LOAD_USER_CACHE:

		users, err := models.Users().All(ctx, tx)

		if err != nil {
			return err
		}

		channelID, err := utility.GetChannelIDByName(ctx, tx, args[0])

		if err != nil {
			return err
		}

		userNameReplacer = utility.UserNameReplacer(users)

		count, err := models.Messages(
			models.MessageWhere.Channel.EQ(channelID),
		).Count(ctx, tx)

		if err != nil {
			return err
		}

		fetchAllMessages, _ := cmd.Flags().GetBool("fetch-all-messages")
		fetchAllMessages = fetchAllMessages || count == 0

		client.KeepFetchingMessages = func(fetchedMessagesCount int, messages []slack.Message) (keepFetching bool) {
			max, _ := cmd.Flags().GetInt("max-fetch-messages")

			if fetchedMessagesCount > max {
				return false
			}

			return fetchAllMessages
		}

		client.KeepFetchingReplies = func(fetchedMessagesCount int, messages []slack.Message) (keepFetching bool) {
			max, _ := cmd.Flags().GetInt("max-fetch-messages")

			if fetchedMessagesCount > max {
				return false
			}

			return fetchAllMessages
		}

		if yes, _ := cmd.Flags().GetBool("offline"); yes {
			goto LOAD_MESSAGE_CACHE
		}
		if err := client.FetchMessages(ctx, tx, channelID); err != nil {
			return err
		}

	LOAD_MESSAGE_CACHE:

		messages, err = models.Messages(
			models.MessageWhere.Channel.EQ(channelID),
			qm.OrderBy(models.MessageColumns.CreatedAt+" DESC"),
		).All(ctx, tx)

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	for _, message := range messages {
		cmd.Printf(
			"@%s %s %s\n",
			userNameReplacer.Replace(message.User),
			utility.MessageReplacer().Replace(userNameReplacer.Replace(message.Text)),
			utility.Ago(message.CreatedAt),
		)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(messageCommand)
	messageCommand.Flags().BoolP("fetch-all-messages", "a", false, "fetch all messages")
	messageCommand.Flags().IntP("max-fetch-messages", "m", 20, "quit fetching when reached this value")
}
