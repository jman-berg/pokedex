package main

import (
	"fmt"
)

func commandHelp(config *Config) error {
	fmt.Print(
		"Welcome to the Pokedex!", 
		"\n",
		"Usage",
		"\n\n\n",
	)
	for k,v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

