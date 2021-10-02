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

var archiveCommand = &cobra.Command{
	Use:     "archive",
	Aliases: []string{"a"},
	Short:   "print single message",
	RunE:    archiveCommandRunE,
}

func archiveCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	channelID, timestamp, err := utility.ExtractChannelIDAndTimestamp(args[0])

	if err != nil {
		return err
	}

	var userNameReplacer *strings.Replacer
	var messages []*models.Message

	err = dbm.Transaction(cmd.Context(), func(ctx context.Context, tx boil.ContextTransactor) error {
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

		userNameReplacer = utility.UserNameReplacer(users)

		count, err := models.Messages(
			models.MessageWhere.Channel.EQ(channelID),
		).Count(ctx, tx)

		if err != nil {
			return err
		}

		fetchAllMessages, _ := cmd.Flags().GetBool("fetch-all-archives")
		fetchAllMessages = fetchAllMessages || count == 0

		client.KeepFetchingMessages = func(fetchedMessagesCount int, archives []slack.Message) (keepFetching bool) {
			if fetchAllMessages {
				return true
			}

			max, _ := cmd.Flags().GetInt("max-fetch-archives")

			return fetchedMessagesCount <= max
		}

		client.KeepFetchingReplies = func(fetchedMessagesCount int, archives []slack.Message) (keepFetching bool) {
			if fetchAllMessages {
				return true
			}

			max, _ := cmd.Flags().GetInt("max-fetch-archives")

			return fetchedMessagesCount <= max
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
		elem := strings.Split(message.Timestamp, ".")

		matchSec := strings.Contains(timestamp, elem[0])
		matchNano := strings.Contains(timestamp, elem[1])

		if !matchSec || !matchNano {
			continue
		}

		cmd.Printf(
			"@%s %s %s\n",
			userNameReplacer.Replace(message.User),
			utility.MessageReplacer().Replace(userNameReplacer.Replace(message.Text)),
			utility.Ago(message.CreatedAt),
		)

		return nil
	}

	cmd.Println("not found")

	return nil
}

func init() {
	RootCommand.AddCommand(archiveCommand)
}
