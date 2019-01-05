package main

import (
	"net/http"
)

// BuildResponse ...
func BuildResponse(w http.ResponseWriter, status int, msg []byte) {
	w.WriteHeader(status)
	w.Write(msg)
}

// JSONResponse ...
func JSONResponse(w http.ResponseWriter, status int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	BuildResponse(w, status, msg)
}

// OKReponse ...
func OKReponse(w http.ResponseWriter) {
	msg := []byte("Success")
	BuildResponse(w, http.StatusOK, msg)
}

// InternalErrorReponse ...
func InternalErrorReponse(w http.ResponseWriter) {
	msg := []byte("There was an internal server error")
	BuildResponse(w, http.StatusInternalServerError, msg)
}

// BadRequestResponse ...
func BadRequestResponse(w http.ResponseWriter) {
	msg := []byte("Bad request")
	BuildResponse(w, http.StatusBadRequest, msg)
}
