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
	Tribes     []Tribe  `json:"tribes"`
	Mode       GameMode `json:"mode"`
	GameLeader *Client  `json:"leader"`
}

// NewGameParams ...
func NewGameParams(leader *Client, mode string, tribes []Tribe) GameParameters {
	gp := GameParameters{
		GameLeader: leader,
		Mode:       str2Mode(mode),
		Tribes:     tribes,
	}
	return gp
}
