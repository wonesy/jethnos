package main

// Message ...
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// JoinMessage ...
type JoinMessage struct {
	HubUUID string `json:"hubUUID"`
}

// MarshalJSON
func (m *Message) MarshalJSON() ([]byte, error) {
	return nil, nil
}
