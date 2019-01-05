package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestNewHub ...
func TestNewHub(t *testing.T) {
	h := NewHub()
	if h == nil {
		t.Error("Returned nil when creating new hub")
	}
}

func TestHttpNewHub(t *testing.T) {
	type response struct {
		HubUUID string `json:"uuid"`
	}

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			CreateHubHandler(w, r)
		}))
	defer ts.Close()

	res, _ := http.Get(ts.URL)

	hubResponse := response{}
	err := json.NewDecoder(res.Body).Decode(&hubResponse)
	if err != nil {
		t.Error("Could not decode json response")
	}

	hub, err := GetHubFromUUID(hubResponse.HubUUID)
	if err != nil || hub == nil {
		t.Errorf("Could not find hub with UUID: %v", hubResponse.HubUUID)
	}
}

func TestNewClient(t *testing.T) {
	type response struct {
		ClientID string `json:"uuid"`
	}

	hub := NewHub()
	go hub.runHub()

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ClientWebsocketHandler(w, r)
		}))
	defer ts.Close()

	url := fmt.Sprintf("%v?hubUUID=%v", ts.URL, hub.id)

	if len(hub.clients) != 0 {
		t.Error("Too many clients registered to hub")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// header requirements to upgrade a req to ws
	req.Header.Set("Connection", "upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-Websocket-Version", "13")
	req.Header.Set("Sec-Websocket-Key", "test")

	client.Do(req)

	timeout := time.NewTicker(1000)

	for {
		select {
		case <-timeout.C:
			t.Error("Timeout on registering client to hub")
		default:
		}

		if len(hub.clients) == 1 {
			break
		}
	}

	var spoke *Client
	for k := range hub.clients {
		spoke = k
		break
	}

	if spoke.hub.id != hub.id {
		t.Error("Spoke and hub IDs don't match")
	}
}
