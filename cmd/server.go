package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mbeka02/broadcast_server/pkg/server"
	"github.com/spf13/cobra"
)

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the broadcast server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		srv := server.NewServer(port, &upgrader)
		srv.Run()
	},
}

func init() {
	serverStartCmd.Flags().StringP("port", "p", "3000", "The port that the broadcast server runs on")
	rootCmd.AddCommand(serverStartCmd)
}
