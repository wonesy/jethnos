package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	id   uuid.UUID
	send chan []byte
	hub  *Hub
	conn *websocket.Conn
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// conveniences
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var (
	// ErrorNilValue ...
	ErrorNilValue = errors.New("Invalid value passed in, cannot be nil")
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ClientWebsocketHandler handles websocket requests from the peer.
func ClientWebsocketHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &Client{
		hub:  nil,
		conn: conn,
		send: make(chan []byte, 256),
	}

	// client.hub.Register <- client

	go client.writePipe()
	go client.readPipe()
}

// RegisterWithHub ...
func (c *Client) RegisterWithHub(hubUUID string) {
	hub, err := GetHubFromUUID(hubUUID)
	if err != nil {
		return
	}

	c.hub = hub
	hub.Register <- c
}

func (c *Client) readPipe() {
	// read from the websocket
	defer func() {
		if c.hub != nil {
			c.hub.Unregister <- c
		}
		close(c.send)
	}()

	// setup default config for the socket reads
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	readMsg := Message{}
	joinMsg := JoinMessage{}

	for {
		// read message from socket
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		fmt.Println("Received message from client")

		err = json.Unmarshal(msg, &readMsg)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(readMsg)

		// ensure that we cannot send messages to hub that don't yet exist
		if c.hub == nil && readMsg.Type != "join" {
			fmt.Println("Cannot send commands to the hub before joining one")
			continue
		}

		switch readMsg.Type {
		case "join":
			fmt.Println("Joining hub...")
			err = json.Unmarshal(msg, &joinMsg)
			if err != nil {
				fmt.Println(err)
			}
			c.RegisterWithHub(joinMsg.Data.HubUUID)
		case "chat":
			c.hub.Broadcast <- bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		default:
			//
		}
	}
}

func (c *Client) writePipe() {
	// write to the websocket
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
