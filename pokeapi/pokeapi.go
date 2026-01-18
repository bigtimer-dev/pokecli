package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, []byte, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, nil, fmt.Errorf("error creating request : %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, nil, fmt.Errorf("error getting response: %w", err)
	}

	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, nil, fmt.Errorf("error reading raw data: %w", err)
	}

	var newLocations Locations

	if err = json.Unmarshal(rawData, &newLocations); err != nil {
		return Locations{}, nil, fmt.Errorf("error decoding the response: %w", err)
	}

	return newLocations, rawData, nil
}

func (c *Client) ExploreLocation(mystring string) (PokemonInArea, []byte, error) {
	const baseURL = "https://pokeapi.co/api/v2/location-area"
	url := baseURL + "/" + mystring
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInArea{}, nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInArea{}, nil, fmt.Errorf("error getting response: %w", err)
	}

	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInArea{}, nil, fmt.Errorf("error reading raw data: %w", err)
	}

	var pokemons PokemonInArea

	if err = json.Unmarshal(rawData, &pokemons); err != nil {
		return PokemonInArea{}, nil, fmt.Errorf("error decoding the response: %w", err)
	}

	return pokemons, rawData, nil
}
