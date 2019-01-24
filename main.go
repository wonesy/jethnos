package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

// BaseURL ...
const BaseURL = "localhost:4444"

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// StartServer ...
func StartServer() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// basic, no auth routes
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc).Methods("GET")
	router.HandleFunc("/ws", ClientWebSocketHandler).Methods("GET")
	router.HandleFunc("/game/list", ListGamesHandler).Methods("GET")
	router.HandleFunc("/game/join/{uuid}", JoinGameHandler).Methods("POST")
	router.HandleFunc("/game/new", NewGameHandler).Methods("POST")

	// routes specific for chat
	authRouter := mux.NewRouter()

	// auth negroni just for authRouter
	an := negroni.New(
		c,
		negroni.HandlerFunc(JwtMiddleware.HandlerWithNext), // token auth
		negroni.Wrap(authRouter),
	)

	router.PathPrefix("/").Handler(an)

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	n.Run(BaseURL)
}

func main() {
	StartServer()
}
