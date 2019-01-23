package main

import "strings"

// GameMode ...
type GameMode int

const (
	// DEMOCRACY ...
	DEMOCRACY GameMode = 0

	// DICTATORSHIP ...
	DICTATORSHIP GameMode = 1
)

func str2Mode(strMode string) GameMode {
	if strings.ToLower(strMode) == "dictatorship" {
		return DICTATORSHIP
	}
	return DEMOCRACY
}

// GameParameters ...
type GameParameters struct {
	Tribes     []Tribe
	Mode       GameMode
	GameLeader *Client
}

// NewGameParams ...
func NewGameParams(leader *Client, mode string, tribes []string) GameParameters {
	gp := GameParameters{
		GameLeader: leader,
		Mode:       str2Mode(mode),
		Tribes:     BuildTribeList(tribes),
	}
	return gp
}
