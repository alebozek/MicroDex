package main

import (
	"fmt"

	"alonso.com/pokemon"
)

func main() {
	Poke := pokemon.CreatePokemon("Charmander")
	fmt.Println(Poke.Moves)
}
