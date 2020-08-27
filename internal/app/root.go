package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moutend/slack/internal/api"
	"github.com/moutend/slack/internal/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client *api.Client
	dbm    *database.Manager
)

var RootCommand = &cobra.Command{
	Use:                "slack",
	Short:              "slack - command line slack client",
	PersistentPreRunE:  rootPersistentPreRunE,
	PersistentPostRunE: rootPersistentPostRunE,
}

func rootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	viper.AutomaticEnv()

	if path, _ := cmd.Flags().GetString("config"); path != "" {
		viper.SetConfigFile(path)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	cacheDirectoryPath := filepath.Join(wd, ".slack")

	os.MkdirAll(cacheDirectoryPath, 0755)

	databaseFilePath := filepath.Join(cacheDirectoryPath, "local.db3")
	migrationDirectoryPath := filepath.Join(cacheDirectoryPath, "migrations")

	if _, err := os.Stat(migrationDirectoryPath); err != nil {
		fmt.Println("@@@")
	}

	dbm, err = database.New(databaseFilePath, migrationDirectoryPath)

	if err != nil {
		return err
	}
	if yes, _ := cmd.Flags().GetBool("debug"); yes {
		dbm.SetLogger(cmd.OutOrStdout())
	}

	botToken := viper.GetString("SLACK_BOT_API_TOKEN")
	userToken := viper.GetString("SLACK_USER_API_TOKEN")

	client = api.New(botToken, userToken)

	if yes, _ := cmd.Flags().GetBool("debug"); yes {
		client.SetLogger(cmd.OutOrStdout())
	}

	return nil
}

func rootPersistentPostRunE(cmd *cobra.Command, args []string) error {
	if dbm != nil {
		dbm.Close()
	}

	return nil
}
func init() {
	RootCommand.PersistentFlags().BoolP("debug", "d", false, "enable debug output")
	RootCommand.PersistentFlags().BoolP("skip-fetch", "s", false, "skip fetch")
	RootCommand.PersistentFlags().StringP("config", "c", "", "path to configuration file")
}
