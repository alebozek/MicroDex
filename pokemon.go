package microdex

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/m7shapan/njson"
)

type Pokemon struct {
	Name   string   `njson:"name"`
	ID     int      `njson:"id"`
	Height int      `njson:"height"`
	Weight int      `njson:"weight"`
	Type_1 string   `njson:"types.0.type.name"`
	Type_2 string   `njson:"types.1.type.name"`
	Moves  []string `njson:"moves.#.move.name"`
}

func CreatePokemon(name string) Pokemon {
	pokemon_url := "https://pokeapi.co/api/v2/pokemon/"
	req, _ := http.Get(pokemon_url + strings.ToLower(name))
	//Checks if the pokemon is on the database
	if req.StatusCode != 200 {
		log.Fatalln("Pokemon not found. Type it again without mistakes.")
	}
	defer req.Body.Close()

	pokejson := Pokemon{}

	json, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("An error ocurred. Error output:", err)
	}
	//Passes the values to the Pokemon struct
	if err := njson.Unmarshal(json, &pokejson); err != nil {
		fmt.Println(err)
	}
	//Capitalizing returning values
	pokejson.Name = strings.Title(pokejson.Name)
	pokejson.Type_1 = strings.Title(pokejson.Type_1)
	pokejson.Type_2 = strings.Title(pokejson.Type_2)
	for i := 0; i < len(pokejson.Moves); i++ {
		pokejson.Moves[i] = strings.Title(pokejson.Moves[i])
	}

	return pokejson
}
