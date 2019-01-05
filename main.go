package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseURL ...
const BaseURL = "localhost:4444"

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// StartServer ...
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc)
	router.HandleFunc("/createhub", CreateHubHandler)
	router.HandleFunc("/joinhub", ClientWebsocketHandler)

	err := http.ListenAndServe(BaseURL, router)
	if err != nil {
		return
	}
}

func main() {
	StartServer()
}
