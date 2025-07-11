package main

import (
	"log"

	"github.com/mbeka02/broadcast_server/pkg/client"
	"github.com/spf13/cobra"
)

var clientConnectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the broadcast server",
	Run: func(cmd *cobra.Command, args []string) {
		serverURL, _ := cmd.Flags().GetString("server")

		c := client.NewClient(serverURL)
		log.Printf("connecting to:%s", serverURL)
		if err := c.Connect(); err != nil {
			log.Fatalf("Failed to connect to %s:%v", serverURL, err)
		}
	},
}

func init() {
	clientConnectCommand.Flags().StringP("server", "s", "ws://localhost:3000/ws", "Server URL to connect to")
	rootCmd.AddCommand(clientConnectCommand)
}
