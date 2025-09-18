package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/jeronimoLa/pokedexcli/internal/pokecache"
)

func indexCache(url string, cfg *Config, c *cache.Cache) {
	pokemon_data := LocationAreaResponse{}
	body, ok := c.Get(url)
	if ok {
		fmt.Println("Cache Found: ", url)
		json.Unmarshal(body, &pokemon_data)
		printResults(pokemon_data.Results)
	} else {
		fmt.Println("Cache Not Found: ", url)
		resp, err := http.Get(url)	
		if err != nil {
			fmt.Println(err)
		}	
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		c.Add(url, body)
		
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


func processRequest(direction string, cfg *Config, c *cache.Cache) {
	urlToHit := ""
	if direction == "Next" {
		urlToHit = cfg.Next
	} else {
		urlToHit = cfg.Previous
	}

	indexCache(urlToHit, cfg, c)

}

func printResults(results []LocationAreaResults) {
	for _, p := range results {
		fmt.Println(p.Name)
	}
}


func processExplore(url string){
	pokemon := PokemonInLocation{}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
	}

	for _, poke := range(pokemon.PokemonEncounters){
		fmt.Println(poke.Pokemon.Name)
	}

}
