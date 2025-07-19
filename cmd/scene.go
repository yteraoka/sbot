/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var sceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "Manage scenes",
	Long:  `Work with SwitchBot scenes. Use subcommands to list or execute scenes.`,
}

func init() {
	rootCmd.AddCommand(sceneCmd)
}