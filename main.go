package main

import (
	"fmt"
	"github.com/404a10/MicroDex"
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
