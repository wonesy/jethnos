package main

import (
	"encoding/json"
	"net/http"
)

// BuildResponse ...
func BuildResponse(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	w.Write(data)
}

// InternalServerErrorResponse ...
func InternalServerErrorResponse(w http.ResponseWriter, msg []byte) {
	BuildResponse(w, http.StatusInternalServerError, msg)
}

// OKResponse ...
func OKResponse(w http.ResponseWriter, msg []byte) {
	BuildResponse(w, http.StatusOK, msg)
}

// BadRequestResponse ...
func BadRequestResponse(w http.ResponseWriter, msg []byte) {
	BuildResponse(w, http.StatusBadRequest, msg)
}

// JSONResponse ...
func JSONResponse(w http.ResponseWriter, code int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		msg := []byte("Could not encode json data")
		InternalServerErrorResponse(w, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
