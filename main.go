package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type requestedPokemon struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Height int    `json:height`
	Weight int    `json:weight`
}

func show(name string, id, height, weight int) {
	fmt.Println("Name:", strings.Title(name))
	fmt.Println("ID:", id)
	fmt.Println("Height:", height*10, "cm")
	fmt.Println("Weight", weight)
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

	if err := json.NewDecoder(r1.Body).Decode(&pokejson); err != nil {
		fmt.Println(err)
	}
	show(pokejson.Name, pokejson.ID, pokejson.Height, pokejson.Weight)
}
