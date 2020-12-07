package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var once sync.Once

// since instance is a DB interface we can swap it out for a mock during testing
var instance DB

// DB - interface for system db
type DB interface {
	getPopulation(string) (int, error)
}

// db - singleton database type
type db struct {
	capitals map[string]int
}

func (db *db) getPopulation(city string) (int, error) {
	res, ok := db.capitals[city]
	var err error
	if !ok {
		err = fmt.Errorf("city: %s, not found", city)
	}
	return res, err
}

// getDB - thread safe, singleton, lazy db initialisation
func getDBInstance() DB {
	once.Do(
		func() {
			caps, err := readDataForDB()
			if err != nil {
				panic(err.Error())
			}
			instance = &db{caps}
		},
	)
	return instance
}

func readDataForDB() (m map[string]int, err error) {
	dat, err := ioutil.ReadFile("./repo/db.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(dat, &m)
	return
}

// GetPopulationOfCities - get the total population of all cities specified in slice
func GetPopulationOfCities(cities []string) (int, error) {
	var total int
	for _, c := range cities {
		pop, err := instance.getPopulation(c)
		if err != nil {
			return 0, err
		}
		total += pop
	}
	return total, nil
}
