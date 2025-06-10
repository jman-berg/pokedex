package main

import (
	"fmt"
)

func commandPokedex(cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}


