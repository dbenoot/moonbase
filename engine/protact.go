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

func look() {
	Output <- lm[player.Coord].Description
	getRoutes(lm[player.Coord])
	getAstroPresent(lm[player.Coord])
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
	}
}

// Functions supporting the actual actions

func getRoutes(v Location) {
	x := v.Coord.X
	y := v.Coord.Y

	if checkCoord(x, y+1) == true {
		Output <- "N: To the north is the " + lm[Coordinates{x, y + 1}].Name + "."
	}

	if checkCoord(x+1, y+1) == true {
		Output <- "NE: To the northeast is the " + lm[Coordinates{x + 1, y + 1}].Name + "."
	}

	if checkCoord(x+1, y) == true {
		Output <- "E: To the east is the " + lm[Coordinates{x + 1, y}].Name + "."
	}

	if checkCoord(x+1, y-1) == true {
		Output <- "SE: To the southeast is the " + lm[Coordinates{x + 1, y - 1}].Name + "."
	}

	if checkCoord(x, y-1) == true {
		Output <- "S: To the south is the " + lm[Coordinates{x, y - 1}].Name + "."
	}

	if checkCoord(x-1, y-1) == true {
		Output <- "SW: To the southwest is the " + lm[Coordinates{x - 1, y - 1}].Name + "."
	}

	if checkCoord(x-1, y) == true {
		Output <- "W: To the west is the " + lm[Coordinates{x - 1, y}].Name + "."
	}

	if checkCoord(x-1, y+1) == true {
		Output <- "NW: To the northwest is the " + lm[Coordinates{x - 1, y + 1}].Name + "."
	}

}

func getAstroPresent(loc Location) {
	for _, v := range npcAstronauts {
		if v.Coord == loc.Coord {
			Output <- v.Name + " is here."
		}
	}
}

func checkCoord(x int, y int) bool {
	if _, ok := lm[Coordinates{x, y}]; ok {
		return true
	}
	return false
}

func GetPlayerStats() string {
	return "HP\t: " + strconv.Itoa(player.hp)
}
