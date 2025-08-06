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

// webhookDeleteCmd represents the delete command
var webhookDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a webhook",
	Long:  `Deletes the webhook for your SwitchBot account.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		c := client.NewClient(token, secret)

		err := c.DeleteWebhook()
		if err != nil {
			return err
		}

		fmt.Println("Successfully deleted webhook.")
		return nil
	},
}

func init() {
	webhookCmd.AddCommand(webhookDeleteCmd)
}
