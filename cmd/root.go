package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "mkctl",
	Short: "CLI tool for managing misskey",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mkctl version: %s\n", version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
