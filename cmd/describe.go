package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe [DEVICE_ID]",
	Short: "Describe a device",
	Long:  `Shows details for a specific device.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		deviceID := args[0]
		c := client.NewClient(token, secret)
		status, err := c.GetDeviceStatus(deviceID)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, status, "", "  "); err != nil {
			return err
		}
		fmt.Println(prettyJSON.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
