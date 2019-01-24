package main

// MaxTribes ...
const MaxTribes = 6

// Tribe ...
type Tribe struct {
	Name  string `json:"name"`
	Power string `json:"power"`
}

var Centaur = Tribe{
	"Centaur",
	"If you place a marker on the board, you may play another band of allies immediately",
}

var Wingfolk = Tribe{
	"Wingfolk",
	"You can place your marker on any kingdom of the board",
}

var Giant = Tribe{
	"Giant",
	"If you play the biggest band with a Giant leader, take control of the Giant Token and score 3 glory points",
}

var Dwarf = Tribe{
	"Dwarf",
	"This band counts as X+1 for end of the age scoring",
}

var Merfolk = Tribe{
	"Merfolk",
	"Advance X spaces on the Merfolk track",
}

var Wizard = Tribe{
	"Wizard",
	"Draw X cards from the deck",
}

var Elf = Tribe{
	"Elf",
	"You may keep up to X cards from your hand",
}

var Halfling = Tribe{
	"Halfling",
	"Don't place a marker on the board",
}

var Minotaur = Tribe{
	"Minotaur",
	"This band counts as X+1 for placing a marker on the board",
}

var Orc = Tribe{
	"Orc",
	"Place a marker on the matching space of the Orc board",
}

var Skeleton = Tribe{
	"Skeleton",
	"Cannot be a leader. Can be used with any band of allies. Must be discarded before end of age scoring",
}

var Troll = Tribe{
	"Troll",
	"Take an unclaimed Troll token with a value up to X",
}

var ErrorTribe = Tribe{
	"Error",
	"Error",
}

func findTribeByName(name string) Tribe {
	for _, tribe := range AllTribes() {
		if tribe.Name == name {
			return tribe
		}
	}
	return ErrorTribe
}

// BuildTribeList ...
func BuildTribeList(nameList []string) []Tribe {
	tribes := []Tribe{}
	for _, name := range nameList {
		tribe := findTribeByName(name)
		tribes = append(tribes, tribe)
	}
	return tribes
}

// AllTribes ...
func AllTribes() []Tribe {
	return []Tribe{
		Centaur,
		Minotaur,
		Wingfolk,
		Giant,
		Dwarf,
		Merfolk,
		Wizard,
		Elf,
		Halfling,
		Orc,
		Skeleton,
		Troll,
	}
}
