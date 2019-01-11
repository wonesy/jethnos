package main

// Message ...
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// JoinMessage ...
type JoinMessage struct {
	Type string          `json:"type"`
	Data JoinMessageData `json:"data"`
}

// JoinMessageData ...
type JoinMessageData struct {
	HubUUID string `json:"hubUUID"`
}
