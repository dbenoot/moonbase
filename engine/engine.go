package engine

import (
	"path/filepath"
	"strconv"

	gameLoop "github.com/kutase/go-gameloop"
)

var glrunning bool
var day, time int
var period, names string
var moonbase Moonbase

// Output exported to the interface
var Output = make(chan string)

// Quit signal exported to the interface
var Quit = make(chan bool)

// Create the PC
var player Astronaut

var locations map[string]Location
var allAstronauts, npcAstronauts []*Astronaut
var lm map[Coordinates]Location
var gl = gameLoop.New(10, func(delta float64) {

	processDateTime()

	player.processAstronaut()

	for _, a := range npcAstronauts {
		a.processAstronaut()
		a.processNPC()
	}
})

// Start function: instances the actors and starts the gameloop
func Start() {
	day = 1
	time = 1

	moonbase = NewBase("Moon Base Alpha", "Research Station", "Self Funded", 10000, 5000, 300, 500, 100, 100)

	// Create some locations

	locations = loadLoc(filepath.Join("engine", "rooms.json"))

	lm = createLocationMap()

	// Create some astronauts

	player = Astronaut{"You", 0, 100, Coordinates{0, 0}, Memory{"working", 1, 100}, []Memory{}}
	a1 := &Astronaut{"Kerbal", 0, 100, Coordinates{4, 0}, Memory{"working", 1, 100}, []Memory{}}
	a2 := &Astronaut{"Leto", 10, 10, Coordinates{2, 0}, Memory{"working", 1, 100}, []Memory{}}

	// allAstronauts = []*Astronaut{player, a1, a2}
	npcAstronauts = []*Astronaut{a1, a2}

	gl.Start()
	glrunning = true

	Output <- "You are in " + moonbase.Name + ", a " + moonbase.Sponsor + " " + moonbase.Government + "."
	Output <- "After 2 years of research on the moon, you are physically unable to return to Earth."
	Output <- "You and your friends have invested most of your assets in buying a small outpost to live in."
}

// Input takes the interface input and processes it. Output is processed by the output channel.
func Input(input string) {
	parse(input)
}

// GetSideBarInfo extracts the sidebar info
func GetSideBarInfo() string {
	return moonbase.Name + "\n" + moonbase.Government + "\n" + moonbase.Sponsor + "\n\n" + strconv.Itoa(moonbase.Money) + "\n" + strconv.Itoa(moonbase.Health) + "\n" + strconv.Itoa(moonbase.Lifesupport) + "\n\n" + strconv.Itoa(moonbase.Water) + "\n" + strconv.Itoa(moonbase.Food) + "\n" + strconv.Itoa(moonbase.Fuel) + "\n\nDay : " + strconv.Itoa(day) + "\nTime : " + period
}

func getTime() string {
	return strconv.Itoa(time)
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
