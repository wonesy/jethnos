package main

import (
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	x, ok := GlobalClientMap[c.UUID]
	if !ok || x != c {
		t.Error("could not find new client in the global map")
	}

	delete(GlobalClientMap, c.UUID)

	if _, ok := GlobalClientMap[c.UUID]; ok {
		t.Error("should have deleted client from global map but it's still there")
	}
}

func TestClientJSON(t *testing.T) {
	c := NewClient()

	_, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
}

/*
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

	url := fmt.Sprintf("%v?hubUUID=%v", ts.URL, hub.UUID)

	if len(hub.Clients) != 0 {
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

		if len(hub.Clients) == 1 {
			break
		}
	}

	var spoke *Client
	for k := range hub.Clients {
		spoke = k
		break
	}

	if spoke.hub.UUID != hub.UUID {
		t.Error("Spoke and hub IDs don't match")
	}
}
*/
