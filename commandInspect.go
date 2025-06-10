package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *Config) error {

	if cfg.parameter == nil {
		return errors.New("inspect requires a pokemon")
	}

	pokemon, exists := cfg.pokedex[*cfg.parameter]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println("  - ",stat.Stat.Name,": ", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println("  - ", t.Type.Name)
	}

	return nil
	

}



