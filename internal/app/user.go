package app

import (
	"context"

	"github.com/moutend/slack/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var userCommand = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u", "users"},
	Short:   "print users",
	RunE:    userCommandRunE,
}

func userCommandRunE(cmd *cobra.Command, args []string) error {
	var users []*models.User

	err := dbm.Transaction(cmd.Context(), func(ctx context.Context, tx boil.ContextTransactor) error {
		if yes, _ := cmd.Flags().GetBool("skip-fetch"); yes {
			goto LOAD_USERS
		}
		if err := client.FetchUsers(ctx, tx); err != nil {
			return err
		}

	LOAD_USERS:

		var err error

		users, err = models.Users(
			qm.OrderBy(models.UserColumns.Name),
		).All(ctx, tx)

		return err
	})
	if err != nil {
		return err
	}
	for _, user := range users {
		cmd.Printf("@%s (%s)\n", user.Name, user.ID)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(userCommand)
}
