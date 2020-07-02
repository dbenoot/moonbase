package engine

import (
	"math/rand"
	"time"
)

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

func (a *Astronaut) move(c Coordinates) {
	if a.checkAP(50) == true {
		if checkCoord(c.X, c.Y) == true {
			a.Coord = c
			Output <- a.Name + " moved to the " + lm[c].Name + "."
			// str := fmt.Sprintf("%#v", a)
			// Output <- str
		} else {
			Output <- a.Name + " cannot move in that direction."
		}
	}
}

func (a *Astronaut) processNPC() {
	if a.ap >= 100 {
		a.decideAction()
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

func (a *Astronaut) decideAction() {
	// test actions: move, think, work, sleep

	switch d := 1; d {
	case 1:
		a.moveNPC()
	case 2:
		a.gotoSleep()
	}
}

func (a *Astronaut) gotoSleep() {}

func (a *Astronaut) moveNPC() {

	r := a.getNPCRoutes()

	// for now have these NPCs wander arounf randomly
	rand.Seed(time.Now().Unix())
	a.move(r[rand.Intn(len(r))])

}

func (a *Astronaut) getNPCRoutes() []Coordinates {
	x := a.Coord.X
	y := a.Coord.Y

	var output []Coordinates

	if checkCoord(x, y+1) == true {
		output = append(output, Coordinates{x, y + 1})
	}

	if checkCoord(x+1, y+1) == true {
		output = append(output, Coordinates{x + 1, y + 1})
	}

	if checkCoord(x+1, y) == true {
		output = append(output, Coordinates{x + 1, y})
	}

	if checkCoord(x+1, y-1) == true {
		output = append(output, Coordinates{x + 1, y - 1})
	}

	if checkCoord(x, y-1) == true {
		output = append(output, Coordinates{x, y - 1})
	}

	if checkCoord(x-1, y-1) == true {
		output = append(output, Coordinates{x - 1, y - 1})
	}

	if checkCoord(x-1, y) == true {
		output = append(output, Coordinates{x - 1, y})
	}

	if checkCoord(x-1, y+1) == true {
		output = append(output, Coordinates{x - 1, y + 1})
	}

	return output

}
