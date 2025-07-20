/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// acCmd represents the ac command
var acCmd = &cobra.Command{
	Use:   "ac",
	Short: "Manage Air Conditioner devices",
	Long:  `Provides subcommands to manage and control Air Conditioner devices.`,
}

func init() {
	rootCmd.AddCommand(acCmd)
}
