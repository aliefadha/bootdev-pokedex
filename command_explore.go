package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("exploring %s .....\n", location.Name)
	fmt.Println("printing pokemon")
	for _, poke := range location.PokemonEncounters {
		fmt.Printf("- %s\n", poke.Pokemon.Name)
	}
	return nil
}
