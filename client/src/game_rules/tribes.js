export class Tribe {
  constructor (name, power) {
    this.name = name
    this.power = power
  }
}

let centaur = new Tribe(
  'Centaur',
  'If you place a marker on the board, you may play another band of allies immediately'
)

let wingfolk = new Tribe(
  'Wingfolk',
  'You can place your marker on any kingdom of the board'
)

let giant = new Tribe(
  'Giant',
  'If you play the biggest band with a Giant leader, take control of the Giant ' +
  'Token and score 3 glory points'
)

let dwarf = new Tribe(
  'Dwarf',
  'This band counts as X+1 for end of the age scoring'
)

let merfolk = new Tribe(
  'Merfolk',
  'Advance X spaces on the Merfolk track'
)

let wizard = new Tribe(
  'Wizard',
  'Draw X cards from the deck'
)

let elf = new Tribe(
  'Elf',
  'You may keep up to X cards from your hand'
)

let halfling = new Tribe(
  'Halfling',
  'Don\'t place a marker on the board'
)

let minotaur = new Tribe(
  'Minotaur',
  'This band counts as X+1 for placing a marker on the board'
)

let orc = new Tribe(
  'Orc',
  'Place a marker on the matching space of the Orc board'
)

let skeleton = new Tribe(
  'Skeleton',
  'Cannot be a leader. Can be used with any band of allies. ' +
  'Must be discarded before end of age scoring'
)

let troll = new Tribe(
  'Troll',
  'Take an unclaimed Troll token with a value up to X'
)

// 100% of all tribes stored in this array
export var JethnosTribes = [
  centaur,
  minotaur,
  wingfolk,
  giant,
  dwarf,
  merfolk,
  wizard,
  elf,
  halfling,
  orc,
  skeleton,
  troll
]
