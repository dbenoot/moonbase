package engine

import (
	"math/rand"
	"time"
)

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

func killAstro(a *Astronaut) {
	for i, aa := range allAstronauts {
		if aa.Name == a.Name {
			allAstronauts[len(allAstronauts)-1], allAstronauts[i] = allAstronauts[i], allAstronauts[len(allAstronauts)-1]
			allAstronauts = allAstronauts[:len(allAstronauts)-1]
		}
	}
}

func (a *Astronaut) decideAction() {
	// test actions: move, think, work, sleep

	switch d := 1; d {
	case 1:
		a.queue = append(a.queue, Action{"move", ""})
	case 2:
		a.gotoSleep()
	}
}

func (a *Astronaut) doAction(act Action) {
	switch act.action {
	case "move":
		a.moveNPC()
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
