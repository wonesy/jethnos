package main

import (
	"testing"
)

func TestGameParams(t *testing.T) {
	tribes := []Tribe{Orc, Skeleton, Centaur, Elf, Merfolk, Wingfolk}
	p := NewGameParams(nil, "dictatorship", tribes)

	if len(p.Tribes) != len(tribes) {
		t.Error("Tribes did not copy properly")
	}

	if p.Mode != DICTATORSHIP {
		t.Error("Wrong game mode")
	}
}

func TestGameState(t *testing.T) {
	tribes := []Tribe{Orc, Skeleton, Centaur, Elf, Merfolk, Wingfolk}
	p := NewGameParams(nil, "dictatorship", tribes)

	gs := CreateGameState(p)

	if gs.Age != 0 {
		t.Error("wrong game age")
	}

	if len(gs.Deck) != (12 * 6) {
		t.Error("Deck size didn't work")
	}

	colormap := make(map[ColorEnum]int)

	for _, card := range gs.Deck {
		val, ok := colormap[card.Color]
		if !ok {
			colormap[card.Color] = 1
		} else {
			colormap[card.Color] = (val + 1)
		}
	}

	for k, v := range colormap {
		if k != BROWN && v != 10 {
			t.Error("Non brown values are incorrect")
		} else if k == BROWN && v != 12 {
			t.Error("Brown values are incorrect")
		}
	}
}
