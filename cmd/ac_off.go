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

// acOffCmd represents the off command
var acOffCmd = &cobra.Command{
	Use:   "off [DEVICE_NAME_OR_ID]",
	Short: "Turn an Air Conditioner off",
	Long:  `Sends the "turnOff" command to a specific Air Conditioner.`,
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

		fmt.Printf("Turned off Air Conditioner %s (%s).\n", nameOrID, deviceID)
		return nil
	},
}

func init() {
	acCmd.AddCommand(acOffCmd)
}
