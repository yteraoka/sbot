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

// sceneListCmd represents the list command
var sceneListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all scenes",
	Long:  `List all scenes registered in your SwitchBot account.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		c := client.NewClient(token, secret)
		scenes, err := c.ListScenes()
		if err != nil {
			return err
		}

		for _, scene := range scenes {
			fmt.Printf("ID: %s, Name: %s\n", scene.ID, scene.Name)
		}

		return nil
	},
}

func init() {
	sceneCmd.AddCommand(sceneListCmd)
}
