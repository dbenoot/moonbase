package engine

type Astronaut struct {
	Name  string
	ap    int
	hp    int
	Coord Coordinates
	Stm   Memory
	Ltm   []Memory
}

type Memory struct {
	memory     string
	value      int // positive or negative memory
	strongness int // How strong is the memory. Higher will create faster and stronger ltm
}

func NewAstronaut(name string, ap int, hp int, coordinates Coordinates, stm Memory, ltm []Memory) Astronaut {
	a := Astronaut{name, ap, hp, coordinates, stm, ltm}
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
	if a.checkAP(50) == true {
		if checkCoord(x, y) == true {
			a.Coord = Coordinates{x, y}
			Output <- "You moved to the " + lm[Coordinates{x, y}].Name + "."
			// str := fmt.Sprintf("%#v", a)
			// Output <- str
		} else {
			Output <- "You cannot move in that direction."
		}
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

func (a *Astronaut) checkAP(i int) bool {
	if a.ap >= i {
		return true
	}
	Output <- "Not enough AP for this action."
	return false
}
