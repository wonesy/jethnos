package main

// MaxTribes ...
const MaxTribes = 6

// Tribe ...
type Tribe struct {
	Name  string `json:"name"`
	Power string `json:"power"`
	Count int    `json:"count"`
}

var Centaur = Tribe{
	"Centaur",
	"If you place a marker on the board, you may play another band of allies immediately",
	12,
}

var Wingfolk = Tribe{
	"Wingfolk",
	"You can place your marker on any kingdom of the board",
	12,
}

var Giant = Tribe{
	"Giant",
	"If you play the biggest band with a Giant leader, take control of the Giant Token and score 3 glory points",
	12,
}

var Dwarf = Tribe{
	"Dwarf",
	"This band counts as X+1 for end of the age scoring",
	12,
}

var Merfolk = Tribe{
	"Merfolk",
	"Advance X spaces on the Merfolk track",
	12,
}

var Wizard = Tribe{
	"Wizard",
	"Draw X cards from the deck",
	12,
}

var Elf = Tribe{
	"Elf",
	"You may keep up to X cards from your hand",
	12,
}

var Halfling = Tribe{
	"Halfling",
	"Don't place a marker on the board",
	24,
}

var Minotaur = Tribe{
	"Minotaur",
	"This band counts as X+1 for placing a marker on the board",
	12,
}

var Orc = Tribe{
	"Orc",
	"Place a marker on the matching space of the Orc board",
	12,
}

var Skeleton = Tribe{
	"Skeleton",
	"Cannot be a leader. Can be used with any band of allies. Must be discarded before end of age scoring",
	12,
}

var Troll = Tribe{
	"Troll",
	"Take an unclaimed Troll token with a value up to X",
	12,
}

var Dragon = Tribe{
	"Dragon",
	"You're getting close to the end of an age...",
	3,
}

var ErrorTribe = Tribe{
	"Error",
	"Error",
	0,
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
