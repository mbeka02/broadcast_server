package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "broadcast-server",
	Short: "A simple  websocket broadcast server",

	Long: `A CLI application for running a WebSocket broadcast server.
    
You can either start a server or connect as a client:
- Use 'start' to run the broadcast server
- Use 'connect' to connect to an existing server`,
	Example: `  broadcast-server start --port 3000
  broadcast-server connect --server ws://localhost:3000/ws`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
