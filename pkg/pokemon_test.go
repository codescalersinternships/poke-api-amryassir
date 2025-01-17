package pkg

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestClientGetPokeByName(t *testing.T) {
	t.Run("Can hit api and return a pokemon", func(t *testing.T) {
		config := LoadConfig()
		client := NewClient(config)
		poke, err := client.GetPokeByName(context.Background(), "pikachu")
		assert.NilError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})
	t.Run("Return an error if pokemon does not exist", func(t *testing.T) {
		config := LoadConfig()
		client := NewClient(config)
		_, err := client.GetPokeByName(context.Background(), "non-existant-pokemon")
		assert.Error(t, err, err.Error())
	})
	t.Run("Can hit locally running server", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"name": "pikachu", "height": 10}`)
			}),
		)
		defer ts.Close()

		config := Config{URL: ts.URL}
		client := NewClient(config)
		poke, err := client.GetPokeByName(context.Background(), "pikachu")
		assert.NilError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
		assert.Equal(t, 10, poke.Height)
	})
}

func TestClientGetAllPokemon(t *testing.T) {
	t.Run("Can hit api and return a list of pokemon", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"results": [{"name": "pikachu", "height": 10}]}`)
			}),
		)
		defer ts.Close()

		config := Config{URL: ts.URL}
		client := NewClient(config)
		pokemonlist, err := client.GetAllPokemon(context.Background())
		assert.NilError(t, err)
		expected := Pokemonlist{
			Results: []Pokemon{
				{Name: "pikachu", Height: 10},
			},
		}
		assert.DeepEqual(t, expected, pokemonlist)
	})

	t.Run("Return an error if status code is not OK", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "not found", http.StatusNotFound)
			}),
		)
		defer ts.Close()

		config := Config{URL: ts.URL}
		client := NewClient(config)
		_, err := client.GetAllPokemon(context.Background())
		assert.Error(t, err, err.Error())
	})
}
