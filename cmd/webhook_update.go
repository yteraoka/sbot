/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

// webhookUpdateCmd represents the update command
var webhookUpdateCmd = &cobra.Command{
	Use:   "update [URL]",
	Short: "Update a webhook",
	Long:  `Updates the webhook for your SwitchBot account.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		webhookURL := args[0]

		c := client.NewClient(token, secret)

		err := c.SetupWebhook(webhookURL)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully updated webhook for URL: %s\n", webhookURL)
		return nil
	},
}

func init() {
	webhookCmd.AddCommand(webhookUpdateCmd)
}
