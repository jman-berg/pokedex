package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(url string) (RespLocationAreas, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}

	if res.StatusCode != http.StatusOK {
		return RespLocationAreas{}, fmt.Errorf("Http request failed with status %s", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	locationAreas := RespLocationAreas{}

	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return RespLocationAreas{}, err
	}

	return locationAreas, nil
}
