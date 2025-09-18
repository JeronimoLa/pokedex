package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/jeronimoLa/pokedexcli/internal/pokecache"
)

func processRequest(direction string, cfg *Config, c *cache.Cache) {
	pokemon_data := apiResponse{}
	urlToHit := ""
	if direction == "Next" {
		urlToHit = cfg.Next
	} else {
		urlToHit = cfg.Previous
	}

	body, ok := c.Get(urlToHit)
	if ok {
		fmt.Println("Cache Found: ", urlToHit)
		json.Unmarshal(body, &pokemon_data)
		printResults(pokemon_data.Results)
	} else {
		fmt.Println("Cache Not Found: ", urlToHit)
		resp, err := http.Get(urlToHit)	
		if err != nil {
			fmt.Println(err)
		}	
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		c.Add(urlToHit, body)
		
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(body, &pokemon_data)
		if err != nil {
			fmt.Println(err)
		}	
		printResults(pokemon_data.Results)
	}
	cfg.Next = pokemon_data.Next
	cfg.Previous = pokemon_data.Previous
}

func printResults(results []PokemonResults) {
	for _, p := range results {
		fmt.Println(p.Name)
	}
}