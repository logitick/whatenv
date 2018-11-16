package main

import (
	"github.com/spf13/cobra"
)

func main()  {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "whatenv",
	Short: "whatenv is an interface for managing releases to environments",
}
