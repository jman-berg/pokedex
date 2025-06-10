package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jman-berg/pokedex/internal/pokecache"
)

func (c *Client) GetLocationAreas(url string, ca *pokecache.Cache) (RespLocationAreas, error) {
	data, exists := ca.Get(url)
	if !exists {
	
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationAreas{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationAreas{}, err
		}

		if res.StatusCode != http.StatusOK {
			return RespLocationAreas{}, fmt.Errorf("Http request failed with status %v", res.StatusCode)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationAreas{}, err
		}
		
		ca.Add(url, data)
		locationAreas := RespLocationAreas{}

		if err := json.Unmarshal(data, &locationAreas); err != nil {
			return RespLocationAreas{}, err
		}

		return locationAreas, nil

	}

	locationAreas := RespLocationAreas{}

	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return RespLocationAreas{}, err
	}

	return locationAreas, nil
}
