package engine

type Astronaut struct {
	Name     string
	Location string
	ap       int
	hp       int
	Coord    Coordinates
}

func NewAstronaut(name string, location string, ap int, hp int, coordinates Coordinates) Astronaut {
	a := Astronaut{name, location, ap, hp, coordinates}
	return a
}

func (a *Astronaut) processAstronaut() {

	a.ap++

	if a.hp <= 0 {
		killAstro(a)
		Output <- "Astronaut" + a.Name + " has died."
	}

}

func (a *Astronaut) move(x int, y int) {
	if checkCoord(x, y) == true {
		a.Location = lm[Coordinates{x, y}].Name
		a.Coord = Coordinates{x, y}
		Output <- "You moved to the " + a.Location + "."
		// str := fmt.Sprintf("%#v", a)
		// Output <- str
	} else {
		Output <- "You cannot move in that direction."
	}
}

func (a *Astronaut) processNPC() {
	if a.ap >= 100 {
		Output <- "Astronaut " + a.Name + " did something."
		a.ap = 0
	}
}

func killAstro(a *Astronaut) {
	for i, aa := range allAstronauts {
		if aa.Name == a.Name {
			allAstronauts[len(allAstronauts)-1], allAstronauts[i] = allAstronauts[i], allAstronauts[len(allAstronauts)-1]
			allAstronauts = allAstronauts[:len(allAstronauts)-1]
		}
	}
}

func GetAstroNames() string {
	astrolist := "ASTRONAUTS\n"
	for _, a := range allAstronauts {
		astrolist = astrolist + "\n" + a.Name
	}
	return astrolist
}
