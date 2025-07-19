/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// bulbCmd represents the bulb command
var bulbCmd = &cobra.Command{
	Use:   "bulb",
	Short: "Manage Color Bulb devices",
	Long:  `Provides subcommands to manage and control Color Bulb devices.`,
}

func init() {
	rootCmd.AddCommand(bulbCmd)
}
