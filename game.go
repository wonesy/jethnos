package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

var (
	// ErrorGameInvalidUUID ...
	ErrorGameInvalidUUID = errors.New("Not a valid UUID")

	// ErrorGameNotFound ...
	ErrorGameNotFound = errors.New("Could not find requested game")
)

// GlobalGameMap will hold a map of all games in play
var GlobalGameMap = make(map[uuid.UUID]*Game)

// Game ...
// holds all of the information about a game including communication channels
type Game struct {
	// This is name given at the time of creation
	Name string

	// This is the unique identifier by which we can refer to the game
	UUID uuid.UUID

	// List of all clients currently registered to the game
	Clients map[*Client]struct{}

	// Permits chat among clients
	ChatBroadcast chan ChatMessage

	// Permits game moves to be registered and broadcast
	MoveBroadcast chan MoveMessage

	// Registers clients to the game
	Register chan *Client

	// Unregisters clients from the game
	Unregister chan *Client

	// Boolean used to determine if the game has already started
	GameStarted bool

	// Used to keep specific parameters for this game
	Parameters GameParameters
}

// MarshalJSON ...
func (g *Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		NumClients  int    `json:"numClients"`
		GameStarted bool   `json:"isStarted"`
	}{
		g.UUID.String(),
		g.Name,
		len(g.Clients),
		g.GameStarted,
	})
}

// GetGameFromUUID ...
func GetGameFromUUID(gameUUID string) (*Game, error) {
	// ensure that the UUID existed in the get params
	if len(gameUUID) == 0 {
		return nil, ErrorGameInvalidUUID
	}

	// ensure it's a valid UUID format
	parsedGameUUID, err := uuid.Parse(gameUUID)
	if err != nil {
		return nil, ErrorGameInvalidUUID
	}

	// ensure that this UUID exists in
	game, ok := GlobalGameMap[parsedGameUUID]
	if !ok {
		return nil, ErrorGameNotFound
	}

	return game, nil
}

// NewGame ...
func NewGame(name string, params GameParameters) *Game {
	uuid, err := uuid.NewUUID()
	if err != nil {
		logger.Error("could not create uuid")
	}

	game := &Game{
		Name:          name,
		UUID:          uuid,
		Clients:       make(map[*Client]struct{}),
		ChatBroadcast: make(chan ChatMessage),
		MoveBroadcast: make(chan MoveMessage),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		GameStarted:   false,
		Parameters:    params,
	}

	logger.Info("created new game: " + game.Name)

	return game
}

func (g *Game) launch() {
	defer func() {
		// close all channels
		close(g.Register)
		close(g.Unregister)
		close(g.ChatBroadcast)
		close(g.MoveBroadcast)
	}()

	for {
		select {
		// Register a new client to the game
		case client := <-g.Register:
			if _, ok := g.Clients[client]; !ok {
				logger.Info("Registering client: " + client.UUID.String())
				g.Clients[client] = struct{}{}
			}

		// Unregister a client from the game
		case client := <-g.Unregister:
			if _, ok := g.Clients[client]; ok {
				logger.Info("Unregistering client: " + client.UUID.String())
				delete(g.Clients, client)
			}

		// Broadcast incoming chat messages
		case msg := <-g.ChatBroadcast:
			for client := range g.Clients {
				select {
				case client.Chat <- msg:
				default:
					// msg could not be sent, remove client
					delete(g.Clients, client)
				}
			}
		}
	}
}

//
// Web handlers for Game
//

// ListGamesHandler ...
func ListGamesHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Game{}
	for _, game := range GlobalGameMap {
		list = append(list, game)
	}
	JSONResponse(w, http.StatusOK, list)
}

// NewGameHandler ...
func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	// read post data
	type postData struct {
		Name   string   `json:"name"`
		Leader string   `json:"uuid"`
		Tribes []string `json:"tribes"`
		Mode   string   `json:"mode"`
	}

	pd := postData{}
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		msg := "Invalid post data"
		logger.Info(msg)
		BadRequestResponse(w, []byte(msg))
		return
	}

	// create a new game
	leader, err := GetClientFromUUID(pd.Leader)
	if err != nil {
		logger.Error(err.Error())
		BadRequestResponse(w, []byte("Could not create game"))
		return
	}

	params := NewGameParams(leader, pd.Mode, pd.Tribes)

	game := NewGame(pd.Name, params)

	// put the new game into the global map
	GlobalGameMap[game.UUID] = game

	// launch the new game
	go game.launch()

	JSONResponse(w, http.StatusOK, game)
}
