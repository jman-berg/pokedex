package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.pokeapiClient.GetLocationAreas(cfg.next, cfg.pokeCache)
	if err != nil {
		return err
	}
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous
	return nil
}


