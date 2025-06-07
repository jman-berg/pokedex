package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)

		if len(cleanedInput) == 0 {
			continue
		}

		firstWord := cleanedInput[0]
		fmt.Print("Your command was: ", firstWord, "\n")
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))

}
