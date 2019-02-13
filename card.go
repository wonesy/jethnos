package main

import "fmt"

// Card ...
type Card struct {
	Tribe Tribe
	Color ColorEnum
}

func addTribeCards(deck *[]Card, color ColorEnum, tribe Tribe) {
	colorSize := tribe.Count / 6
	if tribe.Name == "Skeleton" {
		color = BROWN
	}
	for i := 0; i < colorSize; i++ {
		c := Card{
			Tribe: tribe,
			Color: color,
		}
		*deck = append(*deck, c)
	}
}

// CreateDeck ...
func CreateDeck(tribeList []Tribe) []Card {
	deck := []Card{}
	for _, tribe := range tribeList {
		addTribeCards(&deck, RED, tribe)
		addTribeCards(&deck, BLUE, tribe)
		addTribeCards(&deck, GREEN, tribe)
		addTribeCards(&deck, ORANGE, tribe)
		addTribeCards(&deck, GREY, tribe)
		addTribeCards(&deck, PURPLE, tribe)
	}
	return deck
}

func (c Card) String() string {
	return fmt.Sprintf("%s [%s]", c.Tribe.Name, c.Color.String())
}
