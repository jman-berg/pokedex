package main

import (
	"fmt"
	"errors"
)

func commandMapB(cfg *Config) error {
	if cfg.previous == "" {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
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


