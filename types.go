package main

import cache "github.com/jeronimoLa/pokedexcli/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, c *cache.Cache) error
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
	Next		string
	Previous 	string
	
}

type PokemonInLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
