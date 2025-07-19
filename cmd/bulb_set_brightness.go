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

// bulbSetBrightnessCmd represents the brightness command
var bulbSetBrightnessCmd = &cobra.Command{
	Use:   "brightness [DEVICE_NAME_OR_ID] [LEVEL]",
	Short: "Set the brightness of a Color Bulb",
	Long:  `Sets the brightness for a Color Bulb. LEVEL must be an integer between 1 and 100.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		levelStr := args[1]

		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return fmt.Errorf("invalid brightness level: %s. Must be an integer", levelStr)
		}
		if level < 1 || level > 100 {
			return fmt.Errorf("brightness level must be between 1 and 100, got %d", level)
		}

		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCommand(deviceID, "setBrightness", levelStr)
		if err != nil {
			return err
		}

		fmt.Printf("Set brightness of %s (%s) to %d.\n", nameOrID, deviceID, level)
		return nil
	},
}

func init() {
	bulbSetCmd.AddCommand(bulbSetBrightnessCmd)
}
