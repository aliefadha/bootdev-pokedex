package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		area := Pokemon{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return Pokemon{}, err
		}
		fmt.Println("data from cache:")
		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	area := Pokemon{}
	if err != nil {
		return Pokemon{}, err
	}
	if err := json.Unmarshal(data, &area); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return area, nil
}
