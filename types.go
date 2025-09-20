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
	Pokedex		[]PokemonDetails
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
	BaseExperience int `json:"base_experience"`
	Height    int `json:"height"`
	Name          string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}