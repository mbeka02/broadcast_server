package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	server := newServer(":3000", &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})
	server.Run()
}
