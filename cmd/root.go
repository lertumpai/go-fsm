package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "start application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run the command")
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
