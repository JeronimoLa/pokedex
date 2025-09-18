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
	if cfg.Next == "" {
		cfg.Next = cfg.BaseURL + "location-area/"
	}
	processRequest("Next", cfg, c)
	return nil
}

func commandMapb(cfg *Config, c *cache.Cache, args []string) error {
	if cfg.Previous == "" {
		cfg.Previous = cfg.BaseURL + "location-area/"
	}
	processRequest("Previous", cfg, c)
	return nil
}

func commandExplore(cfg *Config, c *cache.Cache, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing location area")
	}
	if len(args) >1 {
		return fmt.Errorf("too many arguments; expected exactly one location")
	} 

	locationArea := args[0]
	url := cfg.BaseURL + "location-area/" + locationArea
	explore(url)
	//  if err := processExplore(url); err != nil {
    //     return fmt.Errorf("failed to explore %s: %w", url, err)
    // }
	return nil
}

func commandCatch(cfg *Config, c *cache.Cache, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing pokemon name")
	}
	if len(args) >1 {
		return fmt.Errorf("try catching one pokemon at a time")
	} 

	pokemon := args[0]
	url := cfg.BaseURL + "pokemon/" + pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	catch(cfg, url, pokemon)

	fmt.Printf("\n\nPokemon Caught:\n")
	for _, pokemon := range cfg.Pokedex["pokemon_caught"] {
		fmt.Println(pokemon)
	}

	return nil
}