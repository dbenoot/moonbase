package engine

type Astronaut struct {
	Name     string
	Location string
	ap       int
	hp       int
}

func NewAstronaut(name string, location string, ap int, hp int) Astronaut {
	a := Astronaut{name, location, ap, hp}
	return a
}

func (a *Astronaut) processAstronaut() {

	a.ap++
	if a.ap >= 100 {
		Output <- "Astronaut " + a.Name + " did something."
		a.ap = 0
	}

	if a.hp <= 0 {
		killAstro(a)
		Output <- "Astronaut" + a.Name + " has died."
	}

}

func killAstro(a *Astronaut) {
	for i, aa := range astronauts {
		if aa.Name == a.Name {
			astronauts[len(astronauts)-1], astronauts[i] = astronauts[i], astronauts[len(astronauts)-1]
			astronauts = astronauts[:len(astronauts)-1]
		}
	}
}

func GetAstroNames() string {
	astrolist := "ASTRONAUTS\n"
	for _, a := range astronauts {
		astrolist = astrolist + "\n" + a.Name
	}
	return astrolist
}
