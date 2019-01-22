package main

// SocketMessage ...
type SocketMessage struct {
	Type string `json:"type"`
	Data interface{}
}

// ChatMessage ...
type ChatMessage struct {
	Type       string `json:"type"`
	SenderUUID string `json:"uuid"`
	Text       string `json:"text"`
}

// MoveMessage ...
type MoveMessage struct {
	Type string `json:"type"`
}

// JoinGameMessage ...
type JoinGameMessage struct {
	Type     string `json:"type"`
	GameUUID string `json:"uuid"`
}
