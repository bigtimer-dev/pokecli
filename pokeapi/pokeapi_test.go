package pokeapi

import "testing"

func TestExplorePokemon(t *testing.T) {
	client := NewClient()

	data, raw, err := client.ExploreLocation("pastoria-city-area")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(data.PokemonEncounters) == 0 {
		t.Fatal("expected at least 1 pokemon, got 0")
	}

	if len(raw) == 0 {
		t.Fatal("expected raw data, got empty")
	}

	found := false
	for _, p := range data.PokemonEncounters {
		if p.Pokemon.Name == "magikarp" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("expected to find magikarp in pastoria-city-area, but it was not found")
	}
}
