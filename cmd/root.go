/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sbot",
	Short: "A CLI for SwitchBot API",
	Long: `sbot is a command-line interface (CLI) tool for interacting with the SwitchBot API.
It allows you to manage and control your SwitchBot devices from the command line.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// SetVersionInfo sets the version, commit, and date for the application.
func SetVersionInfo(ver, cmt, dt string) {
	version = ver
	commit = cmt
	date = dt
}

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
}
