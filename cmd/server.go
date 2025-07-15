package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mbeka02/broadcast_server/pkg/server"
	"github.com/spf13/cobra"
)

func gracefulShutdown(srv *http.Server, done chan bool) {
	// Create the context that listens for the interrupt signal from the OS
	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-signalCtx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

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
		done := make(chan bool, 1)
		go gracefulShutdown(srv, done)

		log.Println("the server is listening on port" + srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println("http server error:", err)
		}
		<-done
		log.Println("Graceful shutdown is complete")
	},
}

func init() {
	serverStartCmd.Flags().StringP("port", "p", "3000", "The port that the broadcast server runs on")
	rootCmd.AddCommand(serverStartCmd)
}
