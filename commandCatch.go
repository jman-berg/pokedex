package main

import (
	"fmt"
	"time"
	"math/rand"
	"errors"
)

func commandCatch(cfg *Config) error {

	if cfg.parameter == nil {
		return errors.New("catch requires a pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.parameter)

	pokemon, err := cfg.pokeapiClient.GetPokemon("https://pokeapi.co/api/v2/pokemon/"+*cfg.parameter, cfg.pokeCache)
	if err != nil {
		return err
	}

	chance := 50.0 / float64(pokemon.BaseExperience)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	r := rng.Float64()

	if r < chance {
		fmt.Printf("%s was caught!\n", *cfg.parameter)
		cfg.pokedex[*cfg.parameter] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", *cfg.parameter)
	}



	return nil
}
