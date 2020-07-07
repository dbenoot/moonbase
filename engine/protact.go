/*
actions:
- talk
- look
- repair
- move
- take
- sleep
-
*/
package engine

import "strconv"

func look(sub string) {
	if len(sub) == 0 {
		Output <- lm[player.Coord].Description + getRoutes(lm[player.Coord]) + getAstroPresent(lm[player.Coord])
	} else if _, ok := lm[player.Coord].Subsystems[sub]; ok {
		Output <- "The status of the " + sub + " is: " + string(lm[player.Coord].Subsystems[sub]) + "%."
	} else {
		Output <- "There is no " + sub + " in the " + lm[player.Coord].Name + "."
	}
}

func sleep() {
	t = t + 800

	// Make sure your NPCs' time is not stolen from them while you sleep!
	for i := 0; i < 800; i++ {
		for _, a := range npcAstronauts {
			a.processAstronaut()
			a.processNPC()
		}
	}
}

// Temporary debug actions below

func ltmread() {
	for _, v := range npcAstronauts {
		Output <- v.Name
		for i, mem := range v.Ltm {
			Output <- strconv.Itoa(i) + " - " + mem.memory + " - quality : " + strconv.Itoa(mem.quality) + " - persistence : " + strconv.Itoa(mem.persistence)
		}
	}
}

func stmread() {
	for _, v := range npcAstronauts {
		Output <- v.Name
		for i, mem := range v.Stm {
			Output <- strconv.Itoa(i) + " - " + mem.memory + " - quality : " + strconv.Itoa(mem.quality) + " - persistence : " + strconv.Itoa(mem.persistence)
		}
	}
}

func amread() {
	for _, v := range npcAstronauts {
		Output <- v.Name
		Output <- v.Activemem.memory + " - quality : " + strconv.Itoa(v.Activemem.quality) + " - persistence : " + strconv.Itoa(v.Activemem.persistence)
	}
}

func checkap() {
	for _, v := range npcAstronauts {
		Output <- strconv.Itoa(v.ap)
	}
}

func addmem() {
	for _, v := range npcAstronauts {
		v.addActiveMem("newmem", "this is a beautiful new mem", 10)
		Output <- "Added active memory."
	}
}

func clearallactivemem() {
	for _, v := range npcAstronauts {
		v.Activemem.Reset()
	}
}

func (a *Astronaut) clearactivemem() {
	a.Activemem.Reset()
}

var zeroMemory = &Memory{}

func (m *Memory) Reset() {
	*m = *zeroMemory
}

// Functions supporting the actual actions

func getRoutes(v Location) string {
	x := v.Coord.X
	y := v.Coord.Y

	var output string

	if checkCoord(x, y+1) == true {
		output = output + "\nN: To the north is the " + lm[Coordinates{x, y + 1}].Name + "."
	}

	if checkCoord(x+1, y+1) == true {
		output = output + "\nNE: To the northeast is the " + lm[Coordinates{x + 1, y + 1}].Name + "."
	}

	if checkCoord(x+1, y) == true {
		output = output + "\nE: To the east is the " + lm[Coordinates{x + 1, y}].Name + "."
	}

	if checkCoord(x+1, y-1) == true {
		output = output + "\nSE: To the southeast is the " + lm[Coordinates{x + 1, y - 1}].Name + "."
	}

	if checkCoord(x, y-1) == true {
		output = output + "\nS: To the south is the " + lm[Coordinates{x, y - 1}].Name + "."
	}

	if checkCoord(x-1, y-1) == true {
		output = output + "\nSW: To the southwest is the " + lm[Coordinates{x - 1, y - 1}].Name + "."
	}

	if checkCoord(x-1, y) == true {
		output = output + "\nW: To the west is the " + lm[Coordinates{x - 1, y}].Name + "."
	}

	if checkCoord(x-1, y+1) == true {
		output = output + "\nNW: To the northwest is the " + lm[Coordinates{x - 1, y + 1}].Name + "."
	}

	return output

}

func getAstroPresent(loc Location) string {

	var output string

	for _, v := range npcAstronauts {
		if v.Coord == loc.Coord {
			output = output + "\n" + v.Name + " is here. " + v.Name + " is currently doing the task: " + v.action.action + "."
		}
	}

	return output
}

func checkCoord(x int, y int) bool {
	if _, ok := lm[Coordinates{x, y}]; ok {
		return true
	}
	return false
}

func GetProtLoc() Coordinates {
	return player.Coord
}
