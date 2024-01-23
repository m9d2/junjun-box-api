package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use:   "jun jun box api",
		Short: "box",
		Long:  "jun jun box api",
	}
)

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 2020, "Port to listen on")
}
