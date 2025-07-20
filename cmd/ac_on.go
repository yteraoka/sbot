/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yteraoka/sbot/client"
)

var temperature int
var mode string
var fanSpeed string

// acOnCmd represents the on command
var acOnCmd = &cobra.Command{
	Use:   "on [DEVICE_NAME_OR_ID]",
	Short: "Turn an Air Conditioner on with specified settings",
	Long: `Turns on an Air Conditioner using the 'setAll' command.

You must specify the temperature. Mode and fan speed are optional.

Valid modes: auto, cool, dry, fan, heat
Valid fan speeds: auto, low, medium, high`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]

		modeMap := map[string]string{"auto": "1", "cool": "2", "dry": "3", "fan": "4", "heat": "5"}
		fanSpeedMap := map[string]string{"auto": "1", "low": "2", "medium": "3", "high": "4"}

		modeValue, ok := modeMap[strings.ToLower(mode)]
		if !ok {
			return fmt.Errorf("invalid mode: %s. valid modes are auto, cool, dry, fan, heat", mode)
		}

		fanSpeedValue, ok := fanSpeedMap[strings.ToLower(fanSpeed)]
		if !ok {
			return fmt.Errorf("invalid fan speed: %s. valid speeds are auto, low, medium, high", fanSpeed)
		}

		parameter := fmt.Sprintf("%d,%s,%s,on", temperature, modeValue, fanSpeedValue)

		c := client.NewClient(token, secret)

		deviceID, err := c.GetDeviceID(nameOrID)
		if err != nil {
			return err
		}

		err = c.SendCommand(deviceID, "setAll", parameter)
		if err != nil {
			return err
		}

		fmt.Printf("Turned on Air Conditioner %s (%s) with settings: temp=%d, mode=%s, fan=%s.\n", nameOrID, deviceID, temperature, mode, fanSpeed)
		return nil
	},
}

func init() {
	acCmd.AddCommand(acOnCmd)
	acOnCmd.Flags().IntVarP(&temperature, "temperature", "t", 0, "Temperature in Celsius (required)")
	acOnCmd.Flags().StringVarP(&mode, "mode", "m", "auto", "Mode (auto, cool, dry, fan, heat)")
	acOnCmd.Flags().StringVarP(&fanSpeed, "fan-speed", "f", "auto", "Fan speed (auto, low, medium, high)")
	if err := acOnCmd.MarkFlagRequired("temperature"); err != nil {
		panic(err)
	}
}
