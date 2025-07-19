/*
Copyright © 2025 YOUR NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// bulbSetCmd represents the set command
var bulbSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set properties of a Color Bulb",
	Long:  `Set properties like color temperature or color for a Color Bulb.`,
}

func init() {
	bulbCmd.AddCommand(bulbSetCmd)
}
