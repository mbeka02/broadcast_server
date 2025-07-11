package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

type Client struct {
	serverURL string
	conn      *websocket.Conn
}

func NewClient(serverURL string) *Client {
	return &Client{serverURL: serverURL}
}

// creates a new client connection
func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.serverURL, nil)
	if err != nil {
		return fmt.Errorf("client dial error:%v", err)
	}
	c.conn = conn
	// read messages in a separate go-routine
	go c.readMessages()
	c.readUserInput()
	return nil
}

// reads messages from the server
func (c *Client) readMessages() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		fmt.Printf("Received message=>%s\n", msg)
	}
}

func (c *Client) readUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your messages here(Press Enter to send)")
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		switch msg {
		case "":
			continue
		case "/q":
		case "/c":
		case "/quit":
			break
		default:
			err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write error:", err)
				break
			}
		}

	}
	c.conn.Close()
}
