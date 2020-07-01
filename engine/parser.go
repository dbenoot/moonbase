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
		look(subject)
	case "subs":
		Output <- stringify(knownSubjects)
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
		player.move(Coordinates{player.Coord.X, player.Coord.Y + 1})
	case "NE":
		player.move(Coordinates{player.Coord.X + 1, player.Coord.Y + 1})
	case "E":
		player.move(Coordinates{player.Coord.X + 1, player.Coord.Y})
	case "SE":
		player.move(Coordinates{player.Coord.X + 1, player.Coord.Y - 1})
	case "S":
		player.move(Coordinates{player.Coord.X, player.Coord.Y - 1})
	case "SW":
		player.move(Coordinates{player.Coord.X - 1, player.Coord.Y - 1})
	case "W":
		player.move(Coordinates{player.Coord.X - 1, player.Coord.Y})
	case "NW":
		player.move(Coordinates{player.Coord.X - 1, player.Coord.Y + 1})
	default:
		Output <- "Unknown input"
	}
}

func preparse(in string) {

	verblist := []string{
		"look",
		"subs",
		"time",
		"quit",
		"exit",
		"start",
		"sleep",
		"spend",
		"test",
		"ltm",
		"stm",
		"am",
		"ap",
		"addmem",
		"clearmem",
		"N",
		"NE",
		"E",
		"SE",
		"S",
		"SW",
		"W",
		"NW",
	}

	var verb, subject string

	knownVerbs = append(knownVerbs, verblist...)

	Output <- in

	out := strings.Fields(in)
	for _, v := range out {
		if find(knownVerbs, v) == true {
			verb = v
		}
		if find(knownSubjects, v) == true {
			subject = v
		} else {
			subject = ""
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

func stringify(str []string) string {
	var out string

	for _, v := range str {
		out = out + " - " + v
	}

	return out
}
