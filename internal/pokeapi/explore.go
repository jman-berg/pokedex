package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jman-berg/pokedex/internal/pokecache"
)

func (c *Client) GetEncounters(url string, ca *pokecache.Cache) (Encounters, error){
	data, exists := ca.Get(url)
	if !exists {
	
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Encounters{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Encounters{}, err
		}

		if res.StatusCode != http.StatusOK {
			return Encounters{}, fmt.Errorf("Http request failed with status %v", res.StatusCode)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return Encounters{}, err
		}
		
		ca.Add(url, data)
		encounters := Encounters{}

		if err := json.Unmarshal(data, &encounters); err != nil {
			return Encounters{}, err
		}

		return encounters, nil

	}

	encounters := Encounters{}

	if err := json.Unmarshal(data, &encounters); err != nil {
		return Encounters{}, err
	}

	return encounters, nil
	

}
