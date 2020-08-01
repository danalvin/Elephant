package cmd

import (
	"elephant/services"
	"log"

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

	// logging
	log.Println("Starting payment background job ...")

	service := services.NewService()

	service.Logger.Printf("Starting background jobs....\n")

	// List all tasks
	gocron.Every(1).Friday().Do(service.Disburse)

	//run scheduler
	<-gocron.Start()
}
