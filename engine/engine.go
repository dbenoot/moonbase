package engine

var turn int

var moonbase Moonbase

func Start() Moonbase {
	turn = 1

	moonbase = NewBase("Moon Base Alpha", "Research Station", "Self Funded", 10000, 5000, 300, 500, 100, 100)

	return moonbase
}

func Input(input string) (string, int, Moonbase) {

	switch input {
	case "end":
		endTurn()
		return input + " - Turn ended", getTurn(), moonbase
	case "text":
		return getText(), getTurn(), moonbase
	case "spend":
		moonbase.Money = moonbase.Money - 1000
		return input + " - Money spent", getTurn(), moonbase
	default:
		return input + " - Unknown input", getTurn(), moonbase
	}

	return input + " - nothing happened", getTurn(), moonbase
}

func endTurn() {
	calculateTurn()
	turn++
}

func getTurn() int {
	return turn
}

func getText() string {
	return "THIS IS TEXT"
}

// func getStatus() []int {

// }

func calculateTurn() {

}
