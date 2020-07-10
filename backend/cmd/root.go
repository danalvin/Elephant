package cmd

import (
	"elephant/config"
	"elephant/log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "Elephant",
	Aliases: []string{"serve", "s"},
	Short:   "Backend Service for Elephant App",
	Long: `Elephant Microservice for handling automated tasks in paying workers on weekly basis.
		It complements a Django App that runs on the front-end.
	`,
}

// CMD -
type CMD struct {
	Conf   *viper.Viper
	Logger *logrus.Logger
}

var cmd = &CMD{
	Conf:   config.GetConfig(),
	Logger: log.GetLogger(),
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
