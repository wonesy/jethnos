package main

// PlayerState ...
type PlayerState struct {
	TrollTokens []TrollToken
	OrcBoard    OrcBoard
	Hand        []Tribe
}

// HandFull ...
func HandFull(hand []Tribe) bool {
	return len(hand) >= 10
}

// NewPlayerState ...
func NewPlayerState() *PlayerState {
	p := &PlayerState{}
	return p
}
