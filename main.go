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
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// basic, no auth routes
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc)
	router.HandleFunc("/ws", ClientWebsocketHandler)
	router.HandleFunc("/gettoken", GetTokenHandler)
	router.HandleFunc("/createhub", CreateHubHandler) // tmpry for dev work

	// routes specific for chat
	authRouter := mux.NewRouter()
	authRouter.HandleFunc("/listhubs", ListHubHandler)

	// auth negroni just for authRouter
	an := negroni.New(
		c,
		negroni.HandlerFunc(JwtMiddleware.HandlerWithNext), // token auth
		negroni.Wrap(authRouter),
	)

	router.PathPrefix("/").Handler(an)

	n := negroni.New()
	n.Use(c)
	n.UseHandler(router)

	n.Run(BaseURL)
}

func main() {
	StartServer()
}
