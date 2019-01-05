package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var (
	// ErrorHubInvalidUUID ...
	ErrorHubInvalidUUID = errors.New("Not a valid UUID")

	// ErrorHubNotFound ...
	ErrorHubNotFound = errors.New("Could not find requested hub")
)

var globalHubMap = make(map[uuid.UUID]*Hub)

// Hub ...
type Hub struct {
	id         uuid.UUID
	clients    map[*Client]struct{}
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// CreateHubHandler ...
func CreateHubHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		HubUUID string `json:"uuid"`
	}

	if r.Method != "GET" {
		fmt.Printf("Invalid method: %v\n", r.Method)
		BadRequestResponse(w)
		return
	}

	h := NewHub()
	go h.runHub()

	resp := response{h.id.String()}
	data, err := json.Marshal(resp)
	if err != nil {
		InternalErrorReponse(w)
		return
	}

	JSONResponse(w, http.StatusOK, data)
}

// GetHubFromUUID ...
func GetHubFromUUID(hubUUID string) (*Hub, error) {
	// ensure that the UUID existed in the get params
	if len(hubUUID) == 0 {
		return nil, ErrorHubInvalidUUID
	}

	// ensure it's a valid UUID format
	parsedHubUUID, err := uuid.Parse(hubUUID)
	if err != nil {
		return nil, ErrorHubInvalidUUID
	}

	// ensure that this UUID exists in
	hub, ok := globalHubMap[parsedHubUUID]
	if !ok {
		return nil, ErrorHubNotFound
	}

	return hub, nil
}

// NewHub ...
func NewHub() *Hub {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error when creating UUID")
		return nil
	}

	h := &Hub{
		id:         id,
		clients:    make(map[*Client]struct{}),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	globalHubMap[id] = h

	return h
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
			}
		case msg := <-h.broadcast:
			// the hub has received a message to broadcast
			for client := range h.clients {
				select {
				case client.send <- msg:
				default:
					// could not send the message, delete user
					delete(h.clients, client)
				}
			}
		}
	}
}
