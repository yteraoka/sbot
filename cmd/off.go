package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off [DEVICE_NAME_OR_ID]",
	Short: "Turn a device off",
	Long:  `Sends the "turnOff" command to a specific device, specified by its name or ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCommand(deviceID, "turnOff", "default")
		if err != nil {
			return err
		}

		fmt.Printf("Device %s (%s) turned off.\n", nameOrID, deviceID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
