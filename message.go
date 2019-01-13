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

// ChatMessage ...
type ChatMessage struct {
	Type string          `json:"type"`
	Data ChatMessageData `json:"data"`
}

// ChatMessageData ...
type ChatMessageData struct {
	Name       string `json:"name"`
	Text       string `json:"text"`
	SenderUUID string `json:"sender"`
}

// WhoAmIData ...
type WhoAmIData struct {
	Type   string `json:"type"`
	Sender string `json:"uuid"`
}
