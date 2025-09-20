package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

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

func explore(url string) {
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

func catch(cfg *Config, url string) {
	pokemon_details := PokemonDetails{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &pokemon_details)
	if err != nil {
		fmt.Println(err)
	}
	baseExperience := pokemon_details.BaseExperience
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	divisor := baseExperience/50 + 1 // as baseExperience increases, divisor grows, making the event rarer.
	if r.Intn(divisor) == 0 { 
		fmt.Printf("%s was caught!\n", pokemon_details.Name)
		cfg.Pokedex = append(cfg.Pokedex, pokemon_details)
	} else {
		fmt.Printf("%s escaped!\n", pokemon_details.Name)
	}
}

func printPokemon (pokemon PokemonDetails){
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height", pokemon.Height)
	fmt.Println("Weight", pokemon.Weight)
	fmt.Println("Stats")
	for _, val := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types")
	for _, types := range pokemon.Types {
		fmt.Println(" -", types.Type.Name)
	}
}
// catch hippopotas