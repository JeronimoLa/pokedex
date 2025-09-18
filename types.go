package main

import cache "github.com/jeronimoLa/pokedexcli/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, c *cache.Cache) error
}

type PokemonResults struct{
	Name 		string `json:"name"`
	Url			string `json:"url"` 
}

type apiResponse struct {
	Count 		int `json:"count"`
	Next		string `json:"next"`
	Previous 	string `json:"previous"`
	Results 	[]PokemonResults `json:"results"`
}

type Config struct {
	Next		string
	Previous 	string
}
