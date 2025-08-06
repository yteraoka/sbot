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

// webhookGetCmd represents the get command
var webhookGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the webhook URL",
	Long:  `Gets the webhook URL for your SwitchBot account.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		c := client.NewClient(token, secret)

		webhookURL, err := c.GetWebhook()
		if err != nil {
			return err
		}

		fmt.Println(webhookURL)
		return nil
	},
}

func init() {
	webhookCmd.AddCommand(webhookGetCmd)
}
