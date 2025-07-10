package server

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) Run() {
	// start the hub in a separate go-routine
	go s.hub.Run()
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home route")
	})
	router.HandleFunc("/ws", s.handleWs)
	fmt.Println("...server is listening on port", s.Addr)
	err := http.ListenAndServe(s.Addr, router)
	if err != nil {
		log.Fatalf("error , unable to run the server:%v", err)
	}
}

func (s *Server) handleWs(w http.ResponseWriter, r *http.Request) {
	// upgrade HTTP conn to the websocket protocol
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
