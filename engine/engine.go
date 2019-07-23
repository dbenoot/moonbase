package engine

import (
	"strconv"

	gameLoop "github.com/kutase/go-gameloop"
)

var glrunning bool
var turn, day, time int
var period string
var moonbase Moonbase
var Output = make(chan string)

var locations []Location
var astronauts []Astronaut

var gl = gameLoop.New(10, func(delta float64) {

	processDateTime()

	for _, a := range astronauts {
		processAstronaut(a)
	}

})

func Start() {
	turn = 1
	day = 1
	time = 1

	moonbase = NewBase("Moon Base Alpha", "Research Station", "Self Funded", 10000, 5000, 300, 500, 100, 100)

	// Create some locations

	l1 := NewLocation("Laboratory", "This is the laboratory.", []string{"Dormitory", "Airlock"})
	l2 := NewLocation("Dormitory", "This is the dormitory.", []string{"Laboratory"})

	locations = []Location{l1, l2}

	// Create some astronauts

	a1 := NewAstronaut("Kerbal", "Laboratory", 0)
	a2 := NewAstronaut("Leto", "Dormitory", 10)

	astronauts = []Astronaut{a1, a2}

	gl.Start()
	glrunning = true
}

func Input(input string) {
	switch input {
	case "time":
		Output <- getTime()
	case "spend":
		moonbase.Money = moonbase.Money - 1000
		Output <- "Money spent"
	case "pause":
		PauseUnPause()
		Output <- "Pause toggled"
	case "test":
		Output <- "TESTING CHANNEL"
	default:
		Output <- "Unknown input"
	}
}

func GetSideBarInfo() string {
	return moonbase.Name + "\n" + moonbase.Government + "\n" + moonbase.Sponsor + "\n\n" + strconv.Itoa(moonbase.Money) + "\n" + strconv.Itoa(moonbase.Health) + "\n" + strconv.Itoa(moonbase.Lifesupport) + "\n\n" + strconv.Itoa(moonbase.Water) + "\n" + strconv.Itoa(moonbase.Food) + "\n" + strconv.Itoa(moonbase.Fuel) + "\n\nDay : " + strconv.Itoa(day) + "\nTime : " + period
}

func getTime() string {
	return strconv.Itoa(time)
}

func PauseUnPause() {
	if glrunning == true {
		gl.Stop()
		glrunning = false
	} else {
		gl.Start()
		glrunning = true
	}
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

func processAstronaut(a Astronaut) {
	a.AP++
	if a.AP >= 100 {
		Output <- "astronaut did something"
		a.AP = 0
	}
}
