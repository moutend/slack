package app

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var channelCommand = &cobra.Command{
	Use:     "channel",
	Aliases: []string{"c", "channels"},
	Short:   "print channels",
	RunE:    channelCommandRunE,
}

func channelCommandRunE(cmd *cobra.Command, args []string) error {
	var channels []string

	err := dbm.Transaction(cmd.Context(), func(ctx context.Context, tx boil.ContextTransactor) error {
		if yes, _ := cmd.Flags().GetBool("skip-fetch"); yes {
			goto LOAD_CACHE
		}
		if err := client.FetchUsers(ctx, tx); err != nil {
			return err
		}
		if err := client.FetchChannels(ctx, tx); err != nil {
			return err
		}

	LOAD_CACHE:

		query := `
SELECT
  CASE WHEN c.name = ''
  THEN
    CASE WHEN u.name IS NOT NULL THEN '@' || u.name ELSE '' END
  ELSE
    c.name
  END AS name
FROM channels c
LEFT JOIN users u ON u.id = c.user
ORDER BY name
`

		var results []*struct {
			Name string `boil:"name"`
		}

		if err := queries.Raw(query).Bind(ctx, tx, &results); err != nil {
			return err
		}
		for _, result := range results {
			channels = append(channels, result.Name)
		}

		return nil
	})
	if err != nil {
		return err
	}
	for _, channel := range channels {
		cmd.Printf("%s\n", channel)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(channelCommand)
}
