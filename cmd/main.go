package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mbeka02/broadcast_server/pkg/server"
)

func main() {
	server := server.NewServer(":3000", &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})
	server.Run()
}
