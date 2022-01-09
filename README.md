# MicroDex

MicroDex is a golang library based on the RESTful api PokeAPI(https://pokeapi.co/), it offers plenty of information about pokemon, berries, evolutions and so much more about this universe.

In this API I only offer the opportunity to get information about this creatures called pokemon and it's stats.
It's very easy to use, just get the library, import it and create a Pokemon object.

---
## Getting the library
``go get -u github.com/404a10/MicroDex``

---

## Create a Pokemon object
``Charizard := microdex.CreatePokemon("Charizard")``

And that's it, ill show a little demo down here just to show the attributes available in a pokemon object. (I'll try to add everything in the api but, for now I have this).

``
package main

import (
	"fmt"
	microdex "github.com/404a10/MicroDex"
)

func show(pokemon microdex.Pokemon) {
	//We want to print all the information we have from that pokemon
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("ID:", pokemon.ID)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	//If there's no second type, it'll just print the first one
	if pokemon.Type_2 == "" {
		fmt.Println("Type:", pokemon.Type_1)
	} else {
		fmt.Println("Type 1:", pokemon.Type_1)
		fmt.Println("Type 2:", pokemon.Type_2)
	}

	//We'll print the movements as well
	fmt.Println("+--------MOVEMENTS--------+")
	for i := 0; i < len(pokemon.Moves); i++ {
		fmt.Println("-", i, "-", pokemon.Moves[i])
	}
	fmt.Println("+-------------------------+")

}

func main() {
	var p string
	fmt.Println("Enter the name of a pokemon: ")
	fmt.Scanf("%s", &p)
	Pokemon := microdex.CreatePokemon(p)
	show(Pokemon)
}
``

## Output
``
Enter the name of a pokemon: 
Arceus
Name: Arceus
ID: 493
Height: 32
Weight: 3200
Type: Normal
+--------MOVEMENTS--------+
- 0 - Swords-Dance
- 1 - Cut
- 2 - Fly
- 3 - Headbutt
- 4 - Roar
- 5 - Flamethrower
- 6 - Surf
- 7 - Ice-Beam
- 8 - Blizzard
- 9 - Hyper-Beam
- 10 - Seismic-Toss
- 11 - Strength
- 12 - Solar-Beam
- 13 - Thunderbolt
- 14 - Thunder-Wave
- 15 - Thunder
- 16 - Earthquake
- 17 - Toxic
- 18 - Psychic
- 19 - Double-Team
- 20 - Recover
- 21 - Light-Screen
- 22 - Reflect
- 23 - Fire-Blast
- 24 - Waterfall
- 25 - Swift
- 26 - Dream-Eater
- 27 - Flash
- 28 - Rest
- 29 - Rock-Slide
- 30 - Substitute
- 31 - Snore
- 32 - Protect
- 33 - Sludge-Bomb
- 34 - Mud-Slap
- 35 - Perish-Song
- 36 - Icy-Wind
- 37 - Outrage
- 38 - Sandstorm
- 39 - Giga-Drain
- 40 - Endure
- 41 - Swagger
- 42 - Fury-Cutter
- 43 - Sleep-Talk
- 44 - Return
- 45 - Frustration
- 46 - Safeguard
- 47 - Iron-Tail
- 48 - Hidden-Power
- 49 - Twister
- 50 - Rain-Dance
- 51 - Sunny-Day
- 52 - Psych-Up
- 53 - Extreme-Speed
- 54 - Ancient-Power
- 55 - Shadow-Ball
- 56 - Future-Sight
- 57 - Rock-Smash
- 58 - Whirlpool
- 59 - Heat-Wave
- 60 - Hail
- 61 - Will-O-Wisp
- 62 - Facade
- 63 - Trick
- 64 - Magic-Coat
- 65 - Recycle
- 66 - Brick-Break
- 67 - Refresh
- 68 - Secret-Power
- 69 - Dive
- 70 - Hyper-Voice
- 71 - Overheat
- 72 - Rock-Tomb
- 73 - Silver-Wind
- 74 - Cosmic-Power
- 75 - Signal-Beam
- 76 - Bullet-Seed
- 77 - Aerial-Ace
- 78 - Iron-Defense
- 79 - Dragon-Claw
- 80 - Calm-Mind
- 81 - Shock-Wave
- 82 - Water-Pulse
- 83 - Gravity
- 84 - Brine
- 85 - Natural-Gift
- 86 - Tailwind
- 87 - Payback
- 88 - Punishment
- 89 - Last-Resort
- 90 - Poison-Jab
- 91 - Dark-Pulse
- 92 - Aqua-Tail
- 93 - X-Scissor
- 94 - Dragon-Pulse
- 95 - Focus-Blast
- 96 - Energy-Ball
- 97 - Earth-Power
- 98 - Giga-Impact
- 99 - Avalanche
- 100 - Shadow-Claw
- 101 - Zen-Headbutt
- 102 - Flash-Cannon
- 103 - Rock-Climb
- 104 - Defog
- 105 - Trick-Room
- 106 - Draco-Meteor
- 107 - Iron-Head
- 108 - Stone-Edge
- 109 - Stealth-Rock
- 110 - Grass-Knot
- 111 - Judgment
- 112 - Charge-Beam
- 113 - Ominous-Wind
- 114 - Hone-Claws
- 115 - Psyshock
- 116 - Telekinesis
- 117 - Round
- 118 - Echoed-Voice
- 119 - Incinerate
- 120 - Quash
- 121 - Retaliate
- 122 - Bulldoze
- 123 - Work-Up
- 124 - Snarl
- 125 - Confide
+-------------------------+

``
