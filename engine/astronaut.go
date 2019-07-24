package engine

type Astronaut struct {
	Name     string
	Location string
	ap       int
}

func NewAstronaut(name string, location string, ap int) Astronaut {
	a := Astronaut{name, location, ap}
	return a
}

func (a *Astronaut) processAstronaut() {

	a.ap++
	if a.ap >= 100 {
		Output <- "Astronaut " + a.Name + " did something"
		a.ap = 0
	}
}

func GetAstroNames() string {
	astrolist := "ASTRONAUTS\n"
	for _, a := range astronauts {
		astrolist = astrolist + "\n" + a.Name
	}
	return astrolist
}
