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

// Essential gameplay elements

/*
	COLORS
*/

// ColorEnum ...
type ColorEnum int

const (
	// BLUE ...
	BLUE ColorEnum = 0 << iota

	// GREEN ...
	GREEN

	// ORANGE ...
	ORANGE

	// RED ...
	RED

	// PURPLE ...
	PURPLE

	// GREY ...
	GREY

	// BROWN ...
	BROWN
)

/*
	TROLLS
*/

// TrollToken ...
type TrollToken struct {
	Value     int
	IsClaimed bool
}

/*
	ORCS
*/

// OrcBoard ...
type OrcBoard struct {
	Locations [6]OrcLocation
}

// OrcLocation ...
type OrcLocation struct {
	IsUsed bool
	Color  ColorEnum
}

// NewOrcBoard ...
func NewOrcBoard() *OrcBoard {
	ob := &OrcBoard{}
	ob.Locations[0].Color = BLUE
	ob.Locations[0].IsUsed = false
	ob.Locations[1].Color = GREEN
	ob.Locations[1].IsUsed = false
	ob.Locations[2].Color = ORANGE
	ob.Locations[2].IsUsed = false
	ob.Locations[3].Color = PURPLE
	ob.Locations[3].IsUsed = false
	ob.Locations[4].Color = RED
	ob.Locations[4].IsUsed = false
	ob.Locations[5].Color = GREY
	ob.Locations[5].IsUsed = false
	return ob
}
