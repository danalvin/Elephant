package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"elephant/routes"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"serve", "s"},
	Short:   "Elephant Microservice Server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func run() {

	// set routes
	r := routes.Router()

	// server configurations
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", cmd.Conf.GetString("app.port")),
		Handler:        r,
		IdleTimeout:    1 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// server goroutine
	go func() {
		cmd.Logger.Infof("Server up and running on http://localhost%s", s.Addr)
		cmd.Logger.Fatal("Go run go! ", s.ListenAndServe())
	}()

	// for graceful shutdown - channel recieves signal
	// CTRL+C or SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// blocks until a signal is recieved
	sig := <-c
	cmd.Logger.Println("Signal :", sig)

	// context. cancel func complains if ignored.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd.Logger.Println("Server shutting down...")
	s.Shutdown(ctx)
	os.Exit(0)
}
