package main

import "math/rand"

// GameState ...
type GameState struct {
	Age           int
	TrollTokens   [6]TrollToken
	Deck          []Tribe
	DragonIndices [3]int
}

// ShuffleDeck ...
func (gs *GameState) ShuffleDeck() {
	rand.Shuffle(len(gs.Deck), func(i, j int) {
		gs.Deck[i], gs.Deck[j] = gs.Deck[j], gs.Deck[i]
	})

	// Add the 3 dragon cards in
	gs.DragonIndices = [3]int{}
	minDragonIndex := len(gs.Deck) / 2
	maxDragonIndex := len(gs.Deck)
	for i := 0; i < 3; i++ {
		index := rand.Intn(maxDragonIndex-minDragonIndex) + minDragonIndex
		gs.DragonIndices[i] = index
		minDragonIndex = index
	}
}

func (gs *GameState) init(tribeList []Tribe) {
	gs.initDeck(tribeList)
	gs.initTrollTokens()
}

func (gs *GameState) initDeck(tribeList []Tribe) {
	for _, tribe := range tribeList {
		for i := 0; i < tribe.Count; i++ {
			gs.Deck = append(gs.Deck, tribe)
		}
	}
	gs.ShuffleDeck()
}

func (gs *GameState) initTrollTokens() {
	for i, token := range gs.TrollTokens {
		token.Value = 6 - i
		token.IsClaimed = false
	}
}

// CreateGameState ...
func CreateGameState(params GameParameters) GameState {
	gs := GameState{}
	gs.init(params.Tribes)

	return gs
}

//////////////////////////////////////////////////////////////////////////
//
//        Gametime functions
//
//////////////////////////////////////////////////////////////////////////

// AssignPlayerTrollToken ...
func (gs *GameState) AssignPlayerTrollToken(player *PlayerState) {
	for _, token := range gs.TrollTokens {
		if token.IsClaimed == false {
			token.IsClaimed = true
			player.TrollTokens = append(player.TrollTokens, token)
		}
	}
}
