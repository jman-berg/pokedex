package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jman-berg/pokedex/internal/pokecache"
)


func (c *Client) GetPokemon(url string, ca *pokecache.Cache) (Pokemon, error) {
	data, exists := ca.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		if res.StatusCode != http.StatusOK {
			return Pokemon{}, fmt.Errorf("Http request failed with status %v", res.StatusCode)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		
		ca.Add(url, data)
		pokemon := Pokemon{}

		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil

	}

	pokemon := Pokemon{}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
	

}
