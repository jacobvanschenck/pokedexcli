package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex help menu!")
	fmt.Println("\nHere are your available commands:")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
