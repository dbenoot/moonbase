package engine

type Astronaut struct {
	Name      string
	ap        int
	hp        int
	food      int
	rest      int
	Coord     Coordinates
	Activemem Memory
	Stm       []Memory
	Ltm       []Memory
	action    Action
	queue     []Action
}

type Action struct {
	action      string
	actionextra interface{}
}

type Memory struct {
	memory      string
	description string
	quality     int // How positive or negative is the memory. Positive and negative values reflect the positiveness of the memory
	persistence int // Strong memories start with a high value. In Active memory this counts up per turn. In STM and LTM this counts down.
}

func NewAstronaut(name string, ap int, hp int, food int, rest int, coordinates Coordinates, am Memory, stm []Memory, ltm []Memory, action Action, queue []Action) Astronaut {
	a := Astronaut{name, ap, hp, food, rest, coordinates, am, stm, ltm, action, queue}
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

func (a *Astronaut) processNPC() {
	if a.ap >= 100 {
		if len(a.queue) == 0 {
			a.decideAction()
		} else {
			// set first queue item as active action and remove it frmo the queue
			a.action = a.queue[0]
			a.queue = append(a.queue[:0], a.queue[1:]...)

			// perform active action
			a.doAction(a.action)
		}
		a.ap = 0
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
