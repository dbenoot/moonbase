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
	Transitions []string
}

func NewBase(name string, government string, sponsor string, money int, water int, food int, fuel int, lifesupport int, health int) Moonbase {
	m := Moonbase{name, government, sponsor, money, water, food, fuel, lifesupport, health}
	return m
}

func NewLocation(name string, description string, transition []string) Location {
	l := Location{name, description, transition}
	return l
}

func GetLocations() string {
	loclist := "LOCATIONS\n"
	for _, l := range locations {
		loclist = loclist + "\n" + l.Name
	}
	return loclist
}
