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

// tvSetChannelCmd represents the set-channel command
var tvSetChannelCmd = &cobra.Command{
	Use:   "set-channel [DEVICE_NAME_OR_ID] [CHANNEL]",
	Short: "Set the channel of a TV",
	Long:  `Sets the channel for a TV. CHANNEL must be an integer.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		channel := args[1]

		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCommand(deviceID, "SetChannel", channel)
		if err != nil {
			return err
		}

		fmt.Printf("Set channel of %s (%s) to %s.\n", nameOrID, deviceID, channel)
		return nil
	},
}

func init() {
	tvCmd.AddCommand(tvSetChannelCmd)
}

