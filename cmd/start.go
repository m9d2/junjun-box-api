package cmd

import (
	"junjun-box-api/config"
	"junjun-box-api/repository"
	"junjun-box-api/router"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	port = 0

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  "Start the running web server gracefully.",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	slog.Info("Initialize config")
	config.InitConfig()

	slog.Info("Initialize DB")
	repository.Initialize()

	slog.Info("Initialize router")
	r := router.InitRouter()

	slog.Info("Start server", "port", port)
	var err error
	go func() {
		err = r.Run(":" + strconv.Itoa(port))
		if err != nil {
			slog.Error(err.Error())
			return
		} else {
			slog.Info("Service started successfully!")
		}
	}()
	select {}
}

func init() {
	RootCmd.AddCommand(startCmd)
}
