package chat

import (
	"time"

	"github.com/fasthttp/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	Hub        *Hub
	Connection *websocket.Conn
	Send       chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Connection.Close()
	}()
	c.Connection.SetReadLimit(maxMessageSize)
	c.Connection.SetReadDeadline(time.Now().Add(pongWait))
	c.Connection.SetPongHandler(func(string) error { c.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Connection.ReadMessage()
		if err != nil {
			break
		}
		c.Hub.broadcast <- message
	}
}

func (c *Client) writePump() {
}

func PeerChatConnection(hub *Hub, conn *websocket.Conn) {
	client := &Client{Hub: hub, Connection: conn, Send: make(chan []byte, 256)}
	client.Hub.register <- client
	go client.writePump()
	go client.readPump()
}
