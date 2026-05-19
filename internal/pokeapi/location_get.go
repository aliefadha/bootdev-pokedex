package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		area := LocationArea{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return LocationArea{}, err
		}
		fmt.Println("data from cache:")
		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	area := LocationArea{}
	if err != nil {
		return LocationArea{}, err
	}
	if err := json.Unmarshal(data, &area); err != nil {
		return LocationArea{}, err
	}

	return area, nil
}
