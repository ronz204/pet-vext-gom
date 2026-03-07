package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func main() {
	var url string = "https://pokeapi.co/api/v2/pokemon/pikachu"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al consultar", err)
		return
	}

	defer response.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		fmt.Println("Error al decodificar", err)
		return
	}

	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
}
