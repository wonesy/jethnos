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

var (
	// ErrorClientInvalidUUID ...
	ErrorClientInvalidUUID = errors.New("Not a valid UUID")

	// ErrorClientNotFound ...
	ErrorClientNotFound = errors.New("Could not find requested client")
)

var logger = GetLogger()

// GlobalClientMap ...
var GlobalClientMap = make(map[uuid.UUID]*Client)

// Client ...
type Client struct {
	UUID   uuid.UUID
	Socket *websocket.Conn
	Game   *Game
	Chat   chan ChatMessage
}

// MarshalJSON ...
func (c *Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UUID string `json:"uuid"`
	}{
		c.UUID.String(),
	})
}

// // UnmarshalJSON ...
// func (c *Client) UnmarshalJSON(data []byte) error {
// 	aux := struct {
// 		UUID string `json:"uuid"`
// 	}{}

// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}

// 	fmt.Println(aux.UUID)

// 	c, err := GetClientFromUUID(aux.UUID)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		return err
// 	}
// 	return nil
// }

// GetClientFromUUID ...
func GetClientFromUUID(clientUUID string) (*Client, error) {
	// ensure that the UUID existed in the get params
	if len(clientUUID) == 0 {
		return nil, ErrorClientInvalidUUID
	}

	// ensure it's a valid UUID format
	parsedClientUUID, err := uuid.Parse(clientUUID)
	if err != nil {
		return nil, ErrorClientInvalidUUID
	}

	// ensure that this UUID exists in
	client, ok := GlobalClientMap[parsedClientUUID]
	if !ok {
		return nil, ErrorClientNotFound
	}

	return client, nil
}

// NewClient ...
func NewClient() *Client {
	uuid, err := uuid.NewUUID()
	if err != nil {
		logger.Error("could not create uuid")
		return nil
	}

	c := &Client{
		UUID:   uuid,
		Socket: nil,
		Game:   nil,
		Chat:   make(chan ChatMessage),
	}

	return c
}

// RegisterGame ...
func (c *Client) RegisterGame(gameUUID string) error {
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
		delete(GlobalClientMap, c.UUID)
	}()

	// setup default config for the socket reads
	c.Socket.SetReadLimit(maxMessageSize)
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(string) error {
		c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	socketMsg := SocketMessage{}
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
		case "whoami":
			logger.Info("client is requesting a whoami")
			chatMsg.Type = "whoami"
			chatMsg.SenderUUID = c.UUID.String()
			chatMsg.Text = ""
			err := c.Socket.WriteJSON(chatMsg)
			if err != nil {
				logger.Error("could not write whoami response to client")
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

	GlobalClientMap[client.UUID] = client
	logger.Info("Added new client to global list with uuid: " + client.UUID.String())

	client.Socket = conn

	go client.socketReader()
	go client.socketWriter()
}
