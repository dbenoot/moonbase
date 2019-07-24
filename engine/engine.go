package engine

import (
	"strconv"

	gameLoop "github.com/kutase/go-gameloop"
)

var glrunning bool
var day, time int
var period, names string
var moonbase Moonbase
var Output = make(chan string)
var Quit = make(chan bool)

var locations []Location
var astronauts []*Astronaut

var gl = gameLoop.New(10, func(delta float64) {

	processDateTime()

	for _, a := range astronauts {
		a.processAstronaut()
	}
})

func Start() {
	day = 1
	time = 1

	moonbase = NewBase("Moon Base Alpha", "Research Station", "Self Funded", 10000, 5000, 300, 500, 100, 100)

	// Create some locations

	l1 := NewLocation("Laboratory", "This is the laboratory.", []string{"Dormitory", "Airlock"})
	l2 := NewLocation("Dormitory", "This is the dormitory.", []string{"Laboratory"})
	l3 := NewLocation("Airlock", "This is the airlock.", []string{"Laboratory"})

	locations = []Location{l1, l2, l3}

	// Create some astronauts

	a0 := &Astronaut{"You", "Airlock", 0, 100}
	a1 := &Astronaut{"Kerbal", "Laboratory", 0, 100}
	a2 := &Astronaut{"Leto", "Dormitory", 10, 10}

	astronauts = []*Astronaut{a0, a1, a2}

	gl.Start()
	glrunning = true

	// Output <- "You are in " + moonbase.Name + ", a " + moonbase.Sponsor + " " + moonbase.Government + "."
	// Output <- "After 2 years of research on the moon, you are physically unable to return to Earth."
	// Output <- "You and your friends have invested most of your assets in buying a small outpost to live in."
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
	case "exit", "quit":
		Quit <- true
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
		Output <- "Game paused."
	} else {
		gl.Start()
		glrunning = true
		Output <- "Game unpaused."
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
