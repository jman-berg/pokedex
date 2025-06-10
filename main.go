package main

import (
	"time"

	"github.com/jman-berg/pokedex/internal/pokeapi"
	"github.com/jman-berg/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)

	config := &Config{
		pokeapiClient:	pokeClient,
		pokeCache:	pokeCache,
		previous:	"",
		next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		pokedex: map[string]pokeapi.Pokemon{},
		
	}

	startRepl(config)
}


