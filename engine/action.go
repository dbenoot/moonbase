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
	rand.Seed(time.Now().Unix())

	switch rand.Intn(2) {
	case 0:
		a.queue = append(a.queue, Action{"move", ""})
	case 1:
		a.gotoSleep()
	}
}

func (a *Astronaut) doAction(act Action) {

	a.rest = a.rest - 5

	switch act.action {
	case "move":
		r := a.getNPCRoutes()
		// for now have these NPCs wander around randomly
		rand.Seed(time.Now().Unix())
		a.move(r[rand.Intn(len(r))])
	case "sleep":
		a.sleep()
	}
}

func (a *Astronaut) gotoSleep() {

	a.action = Action{"going to sleep", ""}

	// add the sleep action 10 times (time to restore full rest)

	for i := 1; i <= 8; i++ {
		a.queue = append(a.queue, Action{"sleep", ""})
	}
	Output <- a.Name + " is tired and going to sleep."

	// move NPC to the barracks -- this should be changed to any room with a bed subsystem. TODO TO BE UPDATED TO MORE GENERIC BED CONTAINING ROOM SLECETOR

	a.move(Coordinates{2, 1})
}

func (a *Astronaut) sleep() {

	a.addActiveMem("dreaming", "I am having a beautiful dream", 10)

	if a.hp < 100 {
		a.hp = increaseStat(a.hp, 2)
	}
	if a.rest < 100 {
		a.rest = increaseStat(a.rest, 15) // will result in + 10 as doing an action decreases rest already by 5
	}

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

func increaseStat(s int, i int) int {
	s = s + i
	if s > 100 {
		s = 100
	}
	return s
}
