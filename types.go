package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
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
