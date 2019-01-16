
export const tribe = class Tribe {
    constructor (name, power, armyValue, scoreValue) {
        this.name = name
        this.power = power
        this.scoreXMod = scoreXMod
    }
}

let centaur = new Tribe(
    name='Centaur',
    power='If you place a marker on the board, you may play another band of allies immediately'
)

let wingfolk = new Tribe(
    name='Wingfolk',
    power='You can place your marker on any kingdom of the board'
)

let giant = new Tribe(
    name='Giant',
    power='If you play the biggest band with a Giant leader, take control of the Giant '+
        'Token and score 3 glory points'
)

let dwarf = new Tribe(
    name='Dwarf',
    power='This band counts as X+1 for end of the age scoring'
)

let merfolk = new Tribe(
    name='Merfolk',
    power='Advance X spaces on the Merfolk track'
)

let wizard = new Tribe(
    name='Wizard',
    power='Draw X cards from the deck'
)

let elf = new Tribe(
    name='Elf',
    power='You may keep up to X cards from your hand'
)

let halfling = new Tribe(
    name='Halfling',
    power='Don\'t place a marker on the board'
)

let minotaur = new Tribe(
    name='Minotaur',
    power='This band counts as X+1 for placing a marker on the board'
)

let orc = new Tribe(
    name='Orc',
    power='Place a marker on the matching space of the Orc board'
)

let skeleton = new Tribe(
    name='Skeleton',
    power='Cannot be a leader. Can be used with any band of allies. '+
        'Must be discarded before end of age scoring'
)

let troll = new Tribe(
    name='Troll',
    power='Take an unclaimed Troll token with a value up to X'
)

// 100% of all tribes stored in this array
export const JethnosTribes = [
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