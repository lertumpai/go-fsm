package main

import (
	"github.com/spf13/cobra"
	"go-fsm/app"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "start application",
	Run: func(cmd *cobra.Command, args []string) {
		app.App()
	},
}
