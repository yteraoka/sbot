/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of sbot",
	Long:  `All software has versions. This is sbot's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sbot version %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built at: %s\n", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
