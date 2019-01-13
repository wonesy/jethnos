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
	// basic, no auth routes
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc)
	router.HandleFunc("/ws", ClientWebsocketHandler)
	router.HandleFunc("/gettoken", GetTokenHandler)

	// routes specific for chat
	authRouter := mux.NewRouter()
	authRouter.HandleFunc("/createhub", CreateHubHandler)
	authRouter.HandleFunc("/listhubs", ListHubHandler)

	// auth negroni just for authRouter
	an := negroni.New(
		negroni.HandlerFunc(JwtMiddleware.HandlerWithNext), // token auth
		negroni.Wrap(authRouter),                           // include auth routes
	)

	router.PathPrefix("/").Handler(an)

	n := negroni.New()
	n.Use(cors.Default())
	n.UseHandler(router)

	n.Run(BaseURL)
}

func main() {
	StartServer()
}
