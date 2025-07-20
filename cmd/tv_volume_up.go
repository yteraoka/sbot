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

// tvVolumeUpCmd represents the volume-up command
var tvVolumeUpCmd = &cobra.Command{
	Use:   "volume-up [DEVICE_NAME_OR_ID]",
	Short: "Increase the volume of a TV",
	Long:  `Increases the volume of a TV by one step.`,
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

		err = c.SendCommand(deviceID, "volumeAdd", "default")
		if err != nil {
			return err
		}

		fmt.Printf("Increased volume of %s (%s).\n", nameOrID, deviceID)
		return nil
	},
}

func init() {
	tvCmd.AddCommand(tvVolumeUpCmd)
}
