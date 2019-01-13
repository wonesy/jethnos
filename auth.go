package main

import (
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// TODO change to environment
var mySigningKey = []byte("secret")

// GetTokenHandler ...
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	type tokenResponse struct {
		Token string `json:"token"`
	}

	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// set the token claims
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign the token with our secret
	tokenString, _ := token.SignedString(mySigningKey)

	ts := tokenResponse{tokenString}

	// write the value back to the response
	JSONResponse(w, http.StatusOK, ts)
}

// JwtMiddleware ...
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
