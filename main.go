package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/m7shapan/njson"
)

type requestedPokemon struct {
	Name   string   `njson:"name"`
	ID     int      `njson:"id"`
	Height int      `njson:"height"`
	Weight int      `njson:"weight"`
	Type_1 string   `njson:"types.0.type.name"`
	Type_2 string   `njson:"types.1.type.name"`
	Moves  []string `njson:"moves.#.move.name"`
}

func show(pokemon requestedPokemon) {
	fmt.Println("\nName:", strings.Title(pokemon.Name))
	fmt.Println("ID:", pokemon.ID)
	fmt.Println("Height:", pokemon.Height*10, "cm")
	fmt.Println("Weight:", pokemon.Weight, "kg")
	if pokemon.Type_2 == "" {
		fmt.Println("Type:", strings.Title(pokemon.Type_1))
	} else {
		fmt.Println("Type 1:", strings.Title(pokemon.Type_1))
		fmt.Println("Type 2:", strings.Title(pokemon.Type_2))
	}
	fmt.Println("\n##MOVES##")
	for i := 0; i < len(pokemon.Moves); i++ {
		fmt.Printf("Move[%d]: %v \n", i, strings.Title(pokemon.Moves[i]))
	}

}

func main() {
	pokemon_url := "https://pokeapi.co/api/v2/pokemon/"
	var pokemon string
	fmt.Println("Enter the name of the pokemon you want to search: ")
	fmt.Scanf("%s", &pokemon)

	r1, err := http.Get(pokemon_url + strings.ToLower(pokemon))
	if r1.StatusCode != 200 {
		log.Fatalln("Pokemon not found. Type it again without mistakes.")
	}

	if err != nil {
		fmt.Println("An error ocurred. Error output:", err)
	}
	defer r1.Body.Close()
	pokejson := requestedPokemon{}

	json, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Fatal("An error ocurred. Error output:", err)
	}

	if err := njson.Unmarshal(json, &pokejson); err != nil {
		fmt.Println(err)
	}

	show(pokejson)
}
