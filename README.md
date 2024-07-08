# Pokemon API Client

This project provides a Go client for retrieving Pokémon data from an API. The client includes functions to get details about a specific Pokémon by name and to get a list of all available Pokémon. 

## Installation

To use this client, you need to have Go installed on your machine. You can install the package by running:

```bash
go get github.com/codescalersinternships/poke-api-amryassir
```

## Usage
Here's how to use the client in your Go project.
### Initialize a client
``` go
config := LoadConfig()
client := NewClient(config)
```

### GetPokeByName
The GetPokeByName function retrieves data about a specific Pokémon by name.
``` go
pokemon := client.GetPokeByName(context.Background(), "<PokemonName>")
```
### GetAllPokemon
The GetAllPokemon function retrieves a list of all available Pokemon.

``` go
pokemonList := client.GetAllPokemon(context.Background())
```


## Testing
Run the tests using Go's testing package.
```
go test ./...
```
## Contribution
Feel free to open issues or submit pull requests with improvements.

