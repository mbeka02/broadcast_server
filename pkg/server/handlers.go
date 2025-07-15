package server

import (
	"log"
	"net/http"
)

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	// upgrade HTTP connection to the websocket protocol
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(" ws upgrade err:", err)
		return
	}
	// initiate and register the client
	client := &Client{hub: s.hub, conn: conn, send: make(chan []byte, 256)}
	s.hub.register <- client
	// start the go-routines for reading and writing
	go client.writePump()
	go client.readPump()
}
