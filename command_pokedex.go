package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("your pokedex is empty")
		return nil
	}
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" -%s\n", p.Name)
	}

	return nil
}
