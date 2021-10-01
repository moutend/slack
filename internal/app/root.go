package app

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/moutend/slack/internal/api"
	"github.com/moutend/slack/internal/database"
	"github.com/moutend/slack/internal/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client *api.Client
	dbm    *database.Manager

	userName, userID, botName, botID string
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

	db3Path := filepath.Join(cacheDirectoryPath, "cache.db3")

	if err := files.Create(db3Path); err != nil {
		return err
	}

	dbm, err = database.New(db3Path)

	if err != nil {
		return err
	}

	botToken := viper.GetString("SLACK_BOT_API_TOKEN")
	userToken := viper.GetString("SLACK_USER_API_TOKEN")

	client = api.New(botToken, userToken)

	if yes, _ := cmd.Flags().GetBool("debug"); yes {
		client.SetLogger(cmd.OutOrStdout())
		dbm.SetLogger(cmd.OutOrStdout())
	}

	return nil
}

func fetchMigrations(basePath string) error {
	os.MkdirAll(basePath, 0755)

	filenames := []string{
		"0001_db.up.sql",
		"0001_db.down.sql",
	}
	for _, filename := range filenames {
		res, err := http.Get("https://raw.githubusercontent.com/moutend/slack/develop/_migrations/" + filename)

		if err != nil {
			return err
		}

		defer res.Body.Close()

		file, err := os.Create(filepath.Join(basePath, filename))

		if err != nil {
			return err
		}

		defer file.Close()

		if _, err := io.Copy(file, res.Body); err != nil {
			return err
		}
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
