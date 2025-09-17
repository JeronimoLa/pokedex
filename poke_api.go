package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func processRequest(cfg *Config) {
	pokemon_data := apiResponse{}
	resp, err := http.Get(cfg.Next)	
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &pokemon_data)
	if err != nil {
	fmt.Println(err)
	}		

	cfg.Next = pokemon_data.Next
	cfg.Previous = pokemon_data.Previous
	printResults(pokemon_data.Results)
}

func printResults(results []PokemonResults) {
	for _, p := range results {
		fmt.Println(p.Name)
	}
}