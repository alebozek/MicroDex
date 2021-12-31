package main

import (
	"fmt"

	"alonso.com/microdex"
)

func main() {
	Poke := microdex.CreatePokemon("Charmander")
	fmt.Println(Poke.Moves)
}
