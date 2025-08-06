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

// webhookCreateCmd represents the create command
var webhookCreateCmd = &cobra.Command{
	Use:   "create [URL]",
	Short: "Create a webhook",
	Long:  `Creates a webhook for your SwitchBot account.`,
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

		fmt.Printf("Successfully created webhook for URL: %s\n", webhookURL)
		return nil
	},
}

func init() {
	webhookCmd.AddCommand(webhookCreateCmd)
}
