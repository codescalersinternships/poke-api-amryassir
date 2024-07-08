package pkg

// Pokemon - struct represents the deta of a pokemon
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
}

// Pokemonlist - struct represents all pokemons
type Pokemonlist struct {
	Results []Pokemon `json:"results"`
}
