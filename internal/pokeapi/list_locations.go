package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		area := LocationAreas{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return LocationAreas{}, err
		}

		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	area := LocationAreas{}

	if err != nil {
		return LocationAreas{}, err
	}

	if err := json.Unmarshal(data, &area); err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, data)

	return area, nil
}
