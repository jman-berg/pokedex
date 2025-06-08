package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jman-berg/pokedex/internal/pokeapi"
)

type Config struct{
	pokeapiClient pokeapi.Client
	previous string
	next string
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

	}
}






