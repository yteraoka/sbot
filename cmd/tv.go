/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// tvCmd represents the tv command
var tvCmd = &cobra.Command{
	Use:   "tv",
	Short: "Manage TV devices",
	Long:  `Provides subcommands to manage and control TV devices.`,
}

func init() {
	rootCmd.AddCommand(tvCmd)
}
