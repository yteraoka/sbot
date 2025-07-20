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

// runCustomizeCmd represents the run-customize command
var runCustomizeCmd = &cobra.Command{
	Use:   "run-customize [DEVICE_NAME_OR_ID] [BUTTON_NAME]",
	Short: "Execute a custom button on an infrared remote",
	Long:  `Sends a "customize" command to a specific infrared device to execute a custom button (DIY button).`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		buttonName := args[1]

		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCustomizeCommand(deviceID, buttonName)
		if err != nil {
			return err
		}

		fmt.Printf("Executed custom button '%s' on device %s (%s).\n", buttonName, nameOrID, deviceID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCustomizeCmd)
}
