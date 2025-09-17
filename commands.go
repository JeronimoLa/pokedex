package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\t %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *Config) error {
	base := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next == "" {
		cfg.Next = base
	}
	processRequest(cfg)
	return nil
}

func commandMapb(cfg *Config) error {
	return nil
}