/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

// bulbSetColorTemperatureCmd represents the colortemperature command
var bulbSetColorTemperatureCmd = &cobra.Command{
	Use:   "colortemperature [DEVICE_NAME_OR_ID] [KELVIN]",
	Short: "Set the color temperature of a Color Bulb",
	Long:  `Sets the color temperature for a Color Bulb. KELVIN must be an integer between 2700 and 6500.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		kelvinStr := args[1]

		kelvin, err := strconv.Atoi(kelvinStr)
		if err != nil {
			return fmt.Errorf("invalid kelvin value: %s. Must be an integer", kelvinStr)
		}
		if kelvin < 2700 || kelvin > 6500 {
			return fmt.Errorf("kelvin value must be between 2700 and 6500, got %d", kelvin)
		}

		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCommand(deviceID, "setColorTemperature", kelvinStr)
		if err != nil {
			return err
		}

		fmt.Printf("Set color temperature of %s (%s) to %dK.\n", nameOrID, deviceID, kelvin)
		return nil
	},
}

func init() {
	bulbSetCmd.AddCommand(bulbSetColorTemperatureCmd)
}
