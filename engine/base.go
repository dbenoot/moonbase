package engine

import (
	"encoding/json"
	"io/ioutil"
)

type Moonbase struct {
	Name        string
	Government  string
	Sponsor     string
	Money       int
	Water       int // in liters
	Food        int // in kilograms
	Fuel        int // in m³
	Lifesupport int // in %
	Health      int // in %
}

type Location struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Coord       Coordinates    `json:"coord"`
	Subsystems  map[string]int `json:"subsystem"`
}

type Coordinates struct {
	X int
	Y int
}

func NewBase(name string, government string, sponsor string, money int, water int, food int, fuel int, lifesupport int, health int) Moonbase {
	m := Moonbase{name, government, sponsor, money, water, food, fuel, lifesupport, health}
	return m
}

func NewLocation(name string, description string, subsystems map[string]int, coordinates Coordinates) Location {
	l := Location{name, description, coordinates, subsystems}
	return l
}

func createLocationMap() map[Coordinates]Location {
	m := make(map[Coordinates]Location)
	for _, v := range locations {
		m[v.Coord] = v
	}
	return m
}

func loadLoc(f string) map[string]Location {
	var locationMap map[string]Location

	a, err := ioutil.ReadFile(f)
	check(err)

	err = json.Unmarshal(a, &locationMap)
	check(err)

	return locationMap
}
