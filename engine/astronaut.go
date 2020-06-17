package engine

type Astronaut struct {
	Name      string
	ap        int
	hp        int
	Coord     Coordinates
	Activemem Memory
	Stm       []Memory
	Ltm       []Memory
}

type Memory struct {
	memory      string
	description string
	quality     int // How positive or negative is the memory. Positive and negative values reflect the positiveness of the memory
	persistence int // Strong memories start with a high value. In Active memory this counts up per turn. In STM and LTM this counts down.
}

func NewAstronaut(name string, ap int, hp int, coordinates Coordinates, am Memory, stm []Memory, ltm []Memory) Astronaut {
	a := Astronaut{name, ap, hp, coordinates, am, stm, ltm}
	return a
}

func (a *Astronaut) processAstronaut() {

	a.ap++

	// Process the memory items of the astronauts

	a.processMemory()

	// Check if an astronaut died

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

// GetAstroNames creates a list of all of the astronatus.
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

func (a *Astronaut) processMemory() {
	// Then the STM is being forgot at a rate of 2

	a.processStm()

	// Then the LTM is being forgot at a rate of 1

	a.processLtm()

	// every 10 minutes copy the AM to the STM

	if t%600 == 0 {
		a.activeToStm()
	}

	// every hour copy the STM to the LTM

	if t%3600 == 0 {
		a.imprint()
	}
}
