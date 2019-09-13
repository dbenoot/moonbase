package engine

import (
	"math/rand"
	"time"
)

type Astronaut struct {
	Name       string
	ap         int
	hp         int
	Coord      Coordinates
	Activemem  Memory
	Stm        []Memory
	Ltm        []Memory
}

type Memory struct {
	memory      string
	description string
	quality     int // How positive or negative is the memory. Positive and negative values reflect the positiveness of the memory
	timer       int // Strong memories start with a high value. In Active memory this counts up per turn. In STM and LTM this counts down.
}

func NewAstronaut(name string, ap int, hp int, coordinates Coordinates, am Memory, stm []Memory, ltm []Memory) Astronaut {
	a := Astronaut{name, ap, hp, coordinates, am, stm, ltm}
	return a
}

func (a *Astronaut) reminisce() Memory {
	// select a memory from long-term memory and send it to the active memory
	// this should work in such a way that e.g. top 10% of the ltm has 50% chance of returning, while the bottom 50% only 10% for example

	rand.Seed(time.Now().Unix())
	memory := a.Ltm[rand.Intn(len(a.Ltm))]

	return memory
}

func (a *Astronaut) remember() {
	// All short term items are copied to the long term memory with their countdown time.
	// Countdown time is added to the long term memory
	// Countdown time for remembered items must thus be equal to passed time
	// By doing so, very memorable items will remain in the ltm
	// Q: copy every cycle? No, because timer should have time to run out. Every ingame hour or so


}

func (a *Astronaut) processAstronaut() {

	a.ap++

	// Process the memory items of the astronauts
	// First the active memory time is ++

	a.Activemem.timer++

	// Then the STM is being forgot at a rate of 2

	for _, v := range a.Stm {
		v.timer = v.timer-2
	}

	// Then the LTM is being forgot at a rate of 1

	for _, v := range a.Ltm {
		v.timer--
	}

	// every hour copy the STM to the LTM

	if t%60 == 0 {
		a.remember()
	}

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
