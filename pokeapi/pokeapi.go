package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	const baseURL = "https://pokeapi.co/api/v2"
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, fmt.Errorf("error creating request : %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, fmt.Errorf("error getting response: %w", err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var newLocations Locations

	if err = decoder.Decode(&newLocations); err != nil {
		return Locations{}, fmt.Errorf("error decoding the response: %w", err)
	}

	return newLocations, nil
}
