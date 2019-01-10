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
	UUID       uuid.UUID
	Clients    map[*Client]struct{}
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// MarshalJSON ...
func (h *Hub) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UUID       string `json:"uuid"`
		NumClients int    `json:"numClients"`
	}{
		h.UUID.String(),
		len(h.Clients),
	})
}

// CreateHubHandler ...
func CreateHubHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Printf("Invalid method: %v\n", r.Method)
		BadRequestResponse(w, []byte("Invalid method"))
		return
	}

	h := NewHub()
	go h.runHub()

	JSONResponse(w, http.StatusOK, h)
}

// ListHubHandler ...
func ListHubHandler(w http.ResponseWriter, r *http.Request) {
	hubList := []*Hub{}
	for _, h := range globalHubMap {
		hubList = append(hubList, h)
	}
	JSONResponse(w, http.StatusOK, hubList)
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
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error when creating UUID")
		return nil
	}

	h := &Hub{
		UUID:       uuid,
		Clients:    make(map[*Client]struct{}),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	globalHubMap[uuid] = h

	return h
}

func (h *Hub) runHub() {
	for {
		select {
		case client := <-h.Register:
			// We must register a new user to the hub
			h.Clients[client] = struct{}{}
		case client := <-h.Unregister:
			// We must unregister and delete a user from the hub
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
			}
		case msg := <-h.Broadcast:
			// the hub has received a message to broadcast
			for client := range h.Clients {
				select {
				case client.send <- msg:
				default:
					// could not send the message, delete user
					delete(h.Clients, client)
				}
			}
		}
	}
}
