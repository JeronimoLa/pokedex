package main

import (
	"fmt"
	"os"
	"github.com/jeronimoLa/pokedexcli/internal/pokecache"
)

func commandExit(cfg *Config, c *cache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, c *cache.Cache, args []string) error {
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

func commandMap(cfg *Config, c *cache.Cache, args []string) error {
	base := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next == "" {
		cfg.Next = base
	}
	processRequest("Next", cfg, c)
	return nil
}

func commandMapb(cfg *Config, c *cache.Cache, args []string) error {
	base := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous == "" {
		cfg.Previous = base
	}
	processRequest("Previous", cfg, c)
	return nil
}

func commandExplore(cfg *Config, c *cache.Cache, args []string) error {
	base := "https://pokeapi.co/api/v2/location-area/"
	if len(args) == 0 {
		fmt.Println("Location area argument is missing")
	} else if len(args) >1 {
		fmt.Println("Too many arguments; just needs one location")
	} else {
		locationArea := args[0]
		url := base + locationArea
		fmt.Println(url)
		processExplore(url)
	}

	return nil
}