package app

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/moutend/slack/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:               "slack",
	Short:             "slack - command line slack viewer",
	PersistentPreRunE: rootPersistentPreRunE,
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

	botToken := viper.GetString("SLACK_BOT_API_TOKEN")
	userToken := viper.GetString("SLACK_USER_API_TOKEN")

	api = client.New(botToken, userToken, filepath.Join(wd, ".slack"))

	if yes, _ := cmd.Flags().GetBool("debug"); yes {
		api.SetLogger(log.New(os.Stdout, "debug: ", 0))
	}

	return nil
}

func init() {
	rand.Seed(time.Now().Unix())

	RootCommand.PersistentFlags().BoolP("debug", "d", false, "enable debug output")
	RootCommand.PersistentFlags().BoolP("force-fetch", "f", false, "force fetch")
	RootCommand.PersistentFlags().StringP("config", "c", "", "path to configuration file")
}
