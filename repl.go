package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},

		"map": {
			name:        "map",
			description: "List the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View the info of a pokemon you have caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View the list of pokemon you have caught",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
