# MicroDex

MicroDex is a golang library based on the RESTful api PokeAPI(https://pokeapi.co/), it offers plenty of information about pokemon, berries, evolutions and so much more about this universe.

In this API I only offer the opportunity to get information about this creatures called pokemon and it's stats.
It's very easy to use, just get the library, import it and create a Pokemon object.


## Getting the library
``go get -u github.com/404a10/MicroDex``

## Create a Pokemon object
``Charizard := microdex.CreatePokemon("Charizard")``

And that's it, I'll leave a little demo on the demo folder just to show the attributes available in a pokemon object. (I'll try to add everything available about pokemon in the API but, for now I have this).

## List of available attributes
- Name
- ID (The one on the pokedex)
- Height
- Weight
- Types
- Moves

### TO DO
- [ ] Get every attribute on the PokeAPI API about pokemon
- [ ] Add berries to the API
