package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	Addr     string
	Upgrader *websocket.Upgrader
	hub      *Hub
}

func NewServer(addr string, u *websocket.Upgrader) *http.Server {
	NewServer := &Server{
		Addr:     addr,
		Upgrader: u,
		hub:      newHub(),
	}
	// Start the hub in a separate go-routine
	go NewServer.hub.Run()
	return &http.Server{
		Handler:     NewServer.RegisterRoutes(),
		Addr:        ":" + addr,
		IdleTimeout: time.Minute,
	}
}

//
// func (s *Server) Run() {
// 	// start the hub in a separate go-routine
// 	go s.hub.Run()
// 	router := http.NewServeMux()
// 	addr := ":" + s.Addr
// 	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "home route")
// 	})
// 	router.HandleFunc("/ws", s.handleWs)
// 	fmt.Println("...server is listening on port", addr)
// 	err := http.ListenAndServe(addr, router)
// 	if err != nil {
// 		log.Fatalf("error , unable to run the server:%v", err)
// 	}
// }

func (s *Server) RegisterRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home route")
	})
	router.HandleFunc("/ws", s.handleWs)

	return router
}
