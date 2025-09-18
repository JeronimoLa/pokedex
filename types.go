package main

import cache "github.com/jeronimoLa/pokedexcli/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, c *cache.Cache, args []string) error
}

type LocationAreaResults struct{
	Name 		string `json:"name"`
	Url			string `json:"url"` 
}

type LocationAreaResponse struct {
	Count 		int `json:"count"`
	Next		string `json:"next"`
	Previous 	string `json:"previous"`
	Results 	[]LocationAreaResults `json:"results"`
}

type Config struct {
	BaseURL		string
	Next		string
	Previous 	string
	Pokedex		map[string][]string
}

type PokemonInLocation struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name	string `json:"name"`
			URL  	string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonDetails struct {
	BaseExperience	int `json:"base_experience"`
}