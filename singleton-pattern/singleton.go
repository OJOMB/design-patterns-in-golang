package main

import (
	"fmt"

	"github.com/OJOMB/design-patterns-in-golang/singleton-pattern/repo"
)

func main() {
	cities := make([]string, 0)

	cities = append(cities, "Copenhagen")
	cities = append(cities, "London")

	total, _ := repo.GetPopulationOfCities(cities)
	fmt.Println(total)

	cities = append(cities, "Berlin")
	total, _ = repo.GetPopulationOfCities(cities)
	fmt.Println(total)
}
