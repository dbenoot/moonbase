package engine

func parse(input string) {
	switch input {
	case "time":
		Output <- getTime()
	case "sleep":
		sleep()
	case "spend":
		moonbase.Money = moonbase.Money - 1000
		Output <- "Money spent"
	case "pause":
		PauseUnPause()
		Output <- "Pause toggled"
	case "test":
		Output <- "TESTING CHANNEL"
	case "exit", "quit":
		Quit <- true
	case "look":
		look()
	case "N":
		player.move(player.Coord.x, player.Coord.y+1)
	case "NE":
		player.move(player.Coord.x+1, player.Coord.y+1)
	case "E":
		player.move(player.Coord.x+1, player.Coord.y)
	case "SE":
		player.move(player.Coord.x+1, player.Coord.y-1)
	case "S":
		player.move(player.Coord.x, player.Coord.y-1)
	case "SW":
		player.move(player.Coord.x-1, player.Coord.y-1)
	case "W":
		player.move(player.Coord.x-1, player.Coord.y)
	case "NW":
		player.move(player.Coord.x-1, player.Coord.y+1)
	default:
		Output <- "Unknown input"
	}
}
