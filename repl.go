package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jeronimoLa/pokedexcli/internal/pokecache"
)

const pokeAPIBaseURL = "https://pokeapi.co/api/v2/"

func startRepl(){
	c := cache.NewCache(20 * time.Second)
	cfg := &Config{BaseURL: pokeAPIBaseURL}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0{
			continue
		}

		commandName := words[0]
		commandArgs := words[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, c, commandArgs)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	mySlice := strings.Fields(output)
	return mySlice
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations areas in the Pokemon world by increments of 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display pokemon found in a area-location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display pokemon details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of all the names of the Pokemon the user has caught.",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}