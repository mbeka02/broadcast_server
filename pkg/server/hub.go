package server

import "log"

// The Hub is the central coordinator of all the clients. It:
// 1.Manages all client connections
// 2.Handles Message Broadcasting
// 3.Manages registering and unregistering clients.
// NB:It runs in a separate go-routine and uses channels for communication
type Hub struct {
	clients    map[*Client]bool // a map of all connected clients
	broadcast  chan []byte      // inbound messages from the clients
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run() starts the Hub
func (h *Hub) Run() {
	// main hub logic
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("New client connection accepted.Total number of connected clients is:%v\n", len(h.clients))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				// remove client from the map
				delete(h.clients, client)
				// close the client send channel
				close(client.send)
			}
			log.Printf("Client has been disconnected.Total number of clients is:%v\n", len(h.clients))
		case msg := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
