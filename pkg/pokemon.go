package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPokeByName - retrieves a pokemon data using its name
func (c *Client) GetPokeByName(ctx context.Context, pokemonName string) (Pokemon, error) {
	var pokemon Pokemon

	operation := func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", c.config.URL, pokemonName), nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Add("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("response failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		err = json.NewDecoder(resp.Body).Decode(&pokemon)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	if err := Retry(operation); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

// GetAllPokemon - retrieves a list of all pokemons
func (c *Client) GetAllPokemon(ctx context.Context) (Pokemonlist, error) {
	var pokemonlist Pokemonlist
	var err error

	operation := func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.URL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("response failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		err = json.NewDecoder(resp.Body).Decode(&pokemonlist)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	err = Retry(operation)
	if err != nil {
		return Pokemonlist{}, err
	}

	return pokemonlist, nil
}
