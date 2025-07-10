package server

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn // represents the websocket connection
	send chan []byte     // Channel for outbound messages
	hub  *Hub
}

// The readPump() method pumps the message from the web socket connection to the hub
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		c.hub.broadcast <- msg
	}
}

// The writePump() method pumps the message from the hub to the websocket connection

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			// The hub closed the channel
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("write err:", err)
				return
			}
		}
	}
}
