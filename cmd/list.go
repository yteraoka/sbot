package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List devices",
	Long:  `List all devices registered in your SwitchBot account.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		c := client.NewClient(token, secret)
		resp, err := c.ListDevices()
		if err != nil {
			return err
		}

		fmt.Println("Devices:")
		for _, device := range resp.DeviceList {
			fmt.Printf("ID: %s, Name: %s, Type: %s\n", device.ID, device.Name, device.Type)
		}

		if len(resp.InfraredRemoteList) > 0 {
			fmt.Println("\nInfrared Remotes:")
			for _, device := range resp.InfraredRemoteList {
				fmt.Printf("ID: %s, Name: %s, Type: %s\n", device.ID, device.Name, device.Type)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
