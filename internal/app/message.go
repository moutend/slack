package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/moutend/slack/internal/models"
	"github.com/moutend/slack/internal/utility"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var messageCommand = &cobra.Command{
	Use:     "message",
	Aliases: []string{"m", "messages"},
	Short:   "print message",
	RunE:    messageCommandRunE,
}

func messageCommandRunE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return nil
	}

	messageReplacer := strings.NewReplacer(
		"&gt;", ">",
		"&lt;", "<",
		"&amp;", "&",
	)
	var userNameReplacer *strings.Replacer
	var messages []*models.Message

	err := dbm.Transaction(cmd.Context(), func(ctx context.Context, tx boil.ContextTransactor) error {
		var err error

		if yes, _ := cmd.Flags().GetBool("skip-fetch"); yes {
			goto LOAD_CACHE1
		}
		if err := client.FetchUsers(ctx, tx); err != nil {
			return err
		}
		if err := client.FetchChannels(ctx, tx); err != nil {
			return err
		}

	LOAD_CACHE1:

		users, err := models.Users().All(ctx, tx)

		if err != nil {
			return err
		}

		patterns := make([]string, len(users)*4)

		for i, user := range users {
			patterns[i*4] = fmt.Sprintf("<@%s>", user.ID)
			patterns[i*4+1] = fmt.Sprintf("@%s", user.Name)
			patterns[i*4+2] = fmt.Sprintf("%s", user.ID)
			patterns[i*4+3] = fmt.Sprintf("%s", user.Name)
		}

		userNameReplacer = strings.NewReplacer(patterns...)

		query := `
SELECT c.id AS id
FROM channels c
LEFT JOIN users u ON u.id = c.user
WHERE u.name = ? OR c.name = ?
`

		var results []*struct {
			ID string `boil:"id"`
		}

		if err := queries.Raw(query, args[0], args[0]).Bind(ctx, tx, &results); err != nil {
			return err
		}
		if len(results) < 1 {
			return fmt.Errorf("failed to find user or channel: %s", args[0])
		}

		conversationFunc := func(conversationCount int, message slack.Message) bool {
			if yes, _ := cmd.Flags().GetBool("fetch-all-messages"); yes {
				return true
			}

			max, _ := cmd.Flags().GetInt("max-fetch-messages")

			if conversationCount > max {
				return false
			}

			return true
		}
		replyFunc := func(replyCount int, message slack.Message) bool {
			if yes, _ := cmd.Flags().GetBool("fetch-all-messages"); yes {
				return true
			}

			max, _ := cmd.Flags().GetInt("max-fetch-messages")

			if replyCount > max {
				return false
			}

			return true
		}

		if yes, _ := cmd.Flags().GetBool("skip-fetch"); yes {
			goto LOAD_CACHE2
		}
		if err := client.FetchMessages(ctx, tx, results[0].ID, conversationFunc, replyFunc); err != nil {
			return err
		}

	LOAD_CACHE2:

		messages, err = models.Messages(
			models.MessageWhere.Channel.EQ(results[0].ID),
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
			messageReplacer.Replace(userNameReplacer.Replace(message.Text)),
			utility.Ago(message.CreatedAt),
		)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(messageCommand)
	messageCommand.Flags().BoolP("fetch-all-messages", "a", false, "fetch all messages")
	messageCommand.Flags().IntP("max-fetch-messages", "m", 100, "quit fetching when reached this value")
}
