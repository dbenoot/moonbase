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
	case "test":
		Output <- "TESTING CHANNEL"
	case "exit", "quit":
		Quit <- true
	case "look":
		look()
	case "mindread":
		mindread()
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

}
