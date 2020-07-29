package cmd

import (
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var payCMD = &cobra.Command{
	Use:     "pay",
	Aliases: []string{"pay", "p"},
	Short:   "Payment Service CLI",
	Long:    "Payment CLI runs a cron job that periodically disburses amount to wallets",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(payCMD)
}

func start() {
	l.Printf("Starting background jobs....\n")

	// List all tasks
	gocron.Every(1).Minutes().Do()

	//run scheduler
	<-gocron.Start()
}
