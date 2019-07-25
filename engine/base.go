package engine

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
	Name        string
	Description string
	Coord       Coordinates
}

type Coordinates struct {
	x int
	y int
}

func NewBase(name string, government string, sponsor string, money int, water int, food int, fuel int, lifesupport int, health int) Moonbase {
	m := Moonbase{name, government, sponsor, money, water, food, fuel, lifesupport, health}
	return m
}

func NewLocation(name string, description string, coordinates Coordinates) Location {
	l := Location{name, description, coordinates}
	return l
}

func createLocationMap() map[Coordinates]Location {
	m := make(map[Coordinates]Location)
	for _, v := range locations {
		m[v.Coord] = v
	}
	return m
}

func GetLocations() string {
	loclist := "LOCATIONS\n"
	for _, l := range locations {
		loclist = loclist + "\n" + l.Name
	}
	return loclist
}
