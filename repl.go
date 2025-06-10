package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jman-berg/pokedex/internal/pokeapi"
	"github.com/jman-berg/pokedex/internal/pokecache"
)

type Config struct{
	pokeapiClient pokeapi.Client
	pokeCache *pokecache.Cache
	previous string
	next string
	parameter *string
	pokedex map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		if len(cleanedInput) == 2 {
			config.parameter = &cleanedInput[1]
		}
			

		cmd, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(config); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}



type cliCommand struct {
	name		string
	description	string
	callback	func(config *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
		name:		"exit", 
		description:	"Exits the Pokedex",
		callback:	commandExit,
		},
		"help": {
		name:		"help", 
		description:	"Displays a help message",
		callback:	commandHelp,
		},
		"map": {
		name:		"map", 
		description:	"Displays the names of 20 location areas",
		callback:	commandMap,
		},
		"mapb": {
		name:		"mapb", 
		description:	"Displays the names of the previous 20 location areas",
		callback:	commandMapB,
		},
		"explore": {
		name:		"explore", 
		description:	"Lists all the pokemon of a location",
		callback:	commandExplore,
		},
		"catch": {
		name:		"catch", 
		description:	"Catches a given Pokémon and adds it to the user's Pokedex",
		callback:	commandCatch,
		},
		"inspect": {
		name:		"inspect", 
		description:	"See details of a Pokemon you caught",
		callback:	commandInspect,
		},
		"pokedex": {
		name:		"pokedex", 
		description:	"See a list of Pokemon you caught",
		callback:	commandPokedex,
		},

	}
}






