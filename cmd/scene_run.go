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

// sceneRunCmd represents the run command
var sceneRunCmd = &cobra.Command{
	Use:   "run [SCENE_NAME_OR_ID]",
	Short: "Run a scene",
	Long:  `Runs a specific scene, specified by its name or ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("SWITCHBOT_TOKEN")
		secret := os.Getenv("SWITCHBOT_CLIENT_SECRET")
		if token == "" || secret == "" {
			return fmt.Errorf("SWITCHBOT_TOKEN and SWITCHBOT_CLIENT_SECRET must be set")
		}

		nameOrID := args[0]
		c := client.NewClient(token, secret)

		sceneID, err := c.GetSceneID(nameOrID)
		if err != nil {
			return err
		}

		err = c.ExecuteScene(sceneID)
		if err != nil {
			return err
		}

		fmt.Printf("Scene %s (%s) executed.\n", nameOrID, sceneID)
		return nil
	},
}

func init() {
	sceneCmd.AddCommand(sceneRunCmd)
}
