package main

import (
	"github.com/spf13/cobra"
)

var cmdServe = &cobra.Command{
	Use:   "server",
	Short: "Start the webserver",
	Long: `Exposes an api via http`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

