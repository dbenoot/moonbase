package engine

import (
	"strconv"

	gameLoop "github.com/kutase/go-gameloop"
)

var turn, day, time int
var period string
var moonbase Moonbase

//var gl gameLoop

func Start() {
	turn = 1
	day = 1
	time = 1

	moonbase = NewBase("Moon Base Alpha", "Research Station", "Self Funded", 10000, 5000, 300, 500, 100, 100)

	gl := gameLoop.New(10, func(delta float64) {

		processDateTime()

	})

	gl.Start()

}

func Input(input string) string {

	switch input {
	case "text":
		return getText()
	case "spend":
		moonbase.Money = moonbase.Money - 1000
		return input + " - Money spent"
	default:
		return input + " - Unknown input"
	}

	return input + " - nothing happened"
}

func GetSideBarInfo() string {
	return moonbase.Name + "\n" + moonbase.Government + "\n" + moonbase.Sponsor + "\n\n" + strconv.Itoa(moonbase.Money) + "\n" + strconv.Itoa(moonbase.Health) + "\n" + strconv.Itoa(moonbase.Lifesupport) + "\n\n" + strconv.Itoa(moonbase.Water) + "\n" + strconv.Itoa(moonbase.Food) + "\n" + strconv.Itoa(moonbase.Fuel) + "\n\nDay : " + strconv.Itoa(day) + "\nTime : " + strconv.Itoa(time)
}

func getText() string {
	return "THIS IS TEXT"
}

// func getStatus() []int {

// }
func PauseUnPause() {

}

func processDateTime() {
	time++
	if time == 3600 {
		day++
		time = 1
	}
	if time < 1200 {
		period = "Morning"
	} else if time >= 1200 && time < 2400 {
		period = "Afternoon"
	} else {
		period = "Night"
	}
}
