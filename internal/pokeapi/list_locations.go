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

	if err != nil {
		return LocationAreas{}, err
	}

	area := LocationAreas{}

	if err := json.Unmarshal(data, &area); err != nil {
		return LocationAreas{}, err
	}

	return area, nil
}
