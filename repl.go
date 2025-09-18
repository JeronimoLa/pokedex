package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/jeronimoLa/pokedexcli/internal/pokecache"

)

func startRepl(){
	c := cache.NewCache(20 * time.Second)
	cfg := &Config{}
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
				fmt.Println("Error here", err)
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}