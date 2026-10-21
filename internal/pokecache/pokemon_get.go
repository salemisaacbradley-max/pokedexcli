package pokecache

import ("fmt"
		"http"
		"io")

func GetPokemon (pokemon string) error {
	url = "https://pokeapi.co/api/v2/pokemon/" + twine + "/"
	if data, ok := cache.Get(url); ok {
		return data, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error with Map Response: %v\n", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving body: %v\n", err)
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Unexpected Response Code: %v\n", res.StatusCode)
	}
}