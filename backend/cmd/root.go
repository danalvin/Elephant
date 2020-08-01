package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "Elephant",
	Aliases: []string{"serve", "s"},
	Short:   "Backend Service for Elephant App",
	Long: `Elephant Microservice for handling automated tasks in paying workers on weekly basis.
		It complements a Django App that runs on the front-end.
	`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
