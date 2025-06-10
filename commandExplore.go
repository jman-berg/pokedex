package main

import (
	"fmt"
	"errors"
)


func commandExplore(cfg *Config) error {
	if cfg.parameter == nil {
		return errors.New("explore requires a location")
	}
	encounters, err := cfg.pokeapiClient.GetEncounters("https://pokeapi.co/api/v2/location-area/"+*cfg.parameter, cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *cfg.parameter)
	fmt.Println("Found Pokemon") 
	for _, pokemon := range encounters.PokemonEncounters{
		fmt.Println("- ", pokemon.Pokemon.Name)

	}
	return nil 

}
