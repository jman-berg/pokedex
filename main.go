package main

import (
	"time"

	"github.com/jman-berg/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &Config{
		pokeapiClient:	 pokeClient,
		previous:	"",
		next:"https://pokeapi.co/api/v2/location-area/",
		
	}

	startRepl(config)
}


