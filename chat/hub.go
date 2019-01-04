package chatroom

import (
	"fmt"

	"github.com/google/uuid"
)

// Hub ...
type Hub struct {
	id         uuid.UUID
	clients    map[*Client]struct{}
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// NewHub ...
func NewHub() *Hub {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error when creating UUID")
		return nil
	}

	return &Hub{
		id:         id,
		clients:    make(map[*Client]struct{}),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) runHub() {
	for {
		select {
		case client := <-h.register:
			// We must register a new user to the hub
			h.clients[client] = struct{}{}
		case client := <-h.unregister:
			// We must unregister and delete a user from the hub
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case msg := <-h.broadcast:
			// the hub has received a message to broadcast
			for client := range h.clients {
				select {
				case client.send <- msg:
				default:
					// could not send the message, delete user
					delete(h.clients, client)
					close(client.send)
				}
			}
		}
	}
}
