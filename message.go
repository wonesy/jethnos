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
	Name string `json:"name"`
	Text string `json:"text"`
}

// StripLastNewline ...
func (data *ChatMessageData) StripLastNewline() {
	t := data.Text

	if t[len(t)-1] == '\n' {
		data.Text = t[:len(t)-1]
	}
}
