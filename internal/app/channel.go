package app

import (
	"context"

	"github.com/moutend/slack/internal/models"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var channelCommand = &cobra.Command{
	Use:     "channel",
	Aliases: []string{"c", "channels"},
	Short:   "print channels",
	RunE:    channelCommandRunE,
}

type ChannelInfo struct {
	ID     string
	Name   string
	IsUser bool
}

func channelCommandRunE(cmd *cobra.Command, args []string) error {
	result := []ChannelInfo{}

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

		users, err := models.Users().All(ctx, tx)

		if err != nil {
			return err
		}

		channels, err := models.Channels().All(ctx, tx)

		if err != nil {
			return err
		}

		for _, user := range users {
			if user.Name == "" {
				continue
			}

			result = append(result, ChannelInfo{
				ID:     user.ID,
				Name:   "@" + user.Name,
				IsUser: true,
			})
		}
		for _, channel := range channels {
			if channel.Name == "" {
				continue
			}

			result = append(result, ChannelInfo{
				ID:   channel.ID,
				Name: channel.Name,
			})
		}

		return nil
	})
	if err != nil {
		return err
	}
	for _, channelInfo := range result {
		cmd.Println(channelInfo.Name)
	}

	return nil
}

func init() {
	RootCommand.AddCommand(channelCommand)
}
