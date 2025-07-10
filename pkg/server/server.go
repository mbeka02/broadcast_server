package server

import "github.com/gorilla/websocket"

type Server struct {
	Addr     string
	Upgrader *websocket.Upgrader
	hub      *Hub
}

func NewServer(addr string, u *websocket.Upgrader) *Server {
	return &Server{
		Addr:     addr,
		Upgrader: u,
		hub:      newHub(),
	}
}
