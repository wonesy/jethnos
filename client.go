package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client ...
type Client struct {
	UUID   uuid.UUID
	Socket *websocket.Conn
	Game   *Game
	Chat   chan ChatMessage
}

var logger = GetLogger()

// NewClient ...
func NewClient() *Client {
	uuid, err := uuid.NewUUID()
	if err != nil {
		logger.Error("could not create uuid")
		return nil
	}

	return &Client{
		UUID:   uuid,
		Socket: nil,
		Game:   nil,
		Chat:   make(chan ChatMessage),
	}
}

func (c *Client) registerGame(gameUUID string) error {
	game, err := GetGameFromUUID(gameUUID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	c.Game = game
	game.Register <- c

	return nil
}

func (c *Client) socketWriter() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Socket.Close()
	}()

	for {
		select {
		// chat message received, broadcast via Game ChatBroadcast
		case msg, ok := <-c.Chat:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// the game close the channel
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Socket.WriteJSON(msg)
			if err != nil {
				logger.Error("failure to write chat msg to ws")
			}
		case <-ticker.C:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) socketReader() {
	defer func() {
		if c.Game != nil {
			c.Game.Unregister <- c
		}
		close(c.Chat)
	}()

	// setup default config for the socket reads
	c.Socket.SetReadLimit(maxMessageSize)
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(string) error {
		c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	socketMsg := SocketMessage{}
	joinGameMsg := JoinGameMessage{}
	chatMsg := ChatMessage{}

	for {
		// read message from socket
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		logger.Info("received message from client")

		err = json.Unmarshal(msg, &socketMsg)
		if err != nil {
			logger.Error("could not read message in websocket")
			continue
		}

		// do something with this later
		switch socketMsg.Type {
		case "chat":
			logger.Info("sending chat message")
			err = json.Unmarshal(msg, &chatMsg)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			c.Game.ChatBroadcast <- chatMsg
		case "join":
			logger.Info("joining game")
			err = json.Unmarshal(msg, &joinGameMsg)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			err = c.registerGame(joinGameMsg.GameUUID)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
		default:
			logger.Warn("dropping message of unknown type")
		}
	}
}

//
// Web handlers below
//

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

// ClientWebSocketHandler ...
func ClientWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err.Error())
		InternalServerErrorResponse(w, []byte(err.Error()))
		return
	}

	client := NewClient()
	if client == nil {
		msg := "unable to create client"
		logger.Error(msg)
		InternalServerErrorResponse(w, []byte(msg))
		return
	}

	client.Socket = conn

	go client.socketReader()
	go client.socketWriter()
}
