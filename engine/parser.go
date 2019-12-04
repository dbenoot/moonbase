package engine

import (
	"strings"
)

func parse(verb string, subject string) {
	switch verb {
	case "start":
		Start()
	case "time":
		Output <- getTime()
	case "sleep":
		sleep()
	case "spend":
		moonbase.Money = moonbase.Money - 1000
		Output <- "Money spent"
	case "test":
		Output <- "TESTING CHANNEL"
	case "exit", "quit":
		Quit <- true
	case "look":
		look()
	case "ltm":
		ltmread()
	case "stm":
		stmread()
	case "am":
		amread()
	case "ap":
		checkap()
	case "addmem":
		addmem()
	case "clearmem":
		clearallactivemem()
	case "N":
		player.move(player.Coord.X, player.Coord.Y+1)
	case "NE":
		player.move(player.Coord.X+1, player.Coord.Y+1)
	case "E":
		player.move(player.Coord.X+1, player.Coord.Y)
	case "SE":
		player.move(player.Coord.X+1, player.Coord.Y-1)
	case "S":
		player.move(player.Coord.X, player.Coord.Y-1)
	case "SW":
		player.move(player.Coord.X-1, player.Coord.Y-1)
	case "W":
		player.move(player.Coord.X-1, player.Coord.Y)
	case "NW":
		player.move(player.Coord.X-1, player.Coord.Y+1)
	default:
		Output <- "Unknown input"
	}
}

func preparse(in string) {

	var knownVerbs = []string{
		"look",
		"time",
		"quit",
		"exit",
		"start",
	}

	var knownSubjects = []string{
		"wall",
		"Kerbal",
		"Leto",
	}

	var verb, subject string

	Output <- in

	out := strings.Fields(in)
	for _, v := range out {
		if find(knownVerbs, v) == true {
			verb = v
		}
		if find(knownSubjects, v) == true {
			subject = v
		}
	}

	Output <- verb + " : " + subject

	parse(verb, subject)
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
