package engine

type Moonbase struct {
	Name        string
	Government  string
	Sponsor     string
	Money       int
	Water       int // in liters
	Food        int // in kilograms
	Fuel        int // in m³
	Lifesupport int // in %
	Health      int // in %
}

// type Location struct {
// 	Name        string
// 	Description string
// 	Transitions []string
// 	Subsystems  map[string]int
// 	Occupied    string
// 	Specialist  map[string]int
// }

func NewBase(name string, government string, sponsor string, money int, water int, food int, fuel int, lifesupport int, health int) Moonbase {
	m := Moonbase{name, government, sponsor, money, water, food, fuel, lifesupport, health}
	return m
}

// func (s Spaceship) Process(lm map[string]Location, numastro int) Spaceship {

// 	// fuel expenditure based on engine efficiency
// 	s.Fuel = s.Fuel - (100 / float64(lm["Engineering"].Subsystems["engine"]+1)) + 1*(float64(lm["Bridge"].Specialist["Pilot"])/100) + 1*(float64(lm["Engineering"].Specialist["Engineering"])/100) + 1.3*(float64(lm["Outside in space"].Specialist["Engineering"])/100)

// 	// oxygen expenditure based on #astronauts and oxygen reclamantion efficiency
// 	s.Oxygen = s.Oxygen - float64(numastro)*0.5 + float64(lm["Engineering"].Subsystems["oxygen"])/100*1.5 + 1*float64(lm["Specialist"].Subsystems["Maintenance"])/100
// 	if s.Oxygen > 22 {
// 		s.Oxygen = 22 //too much oxygen is regulated automatically
// 	}

// 	// CO2

// 	s.CO2 = int(float64(s.CO2) + float64(numastro)*float64(100) - float64(lm["Engineering"].Subsystems["carbon dioxide filter"])*3.125)
// 	if s.CO2 < 0 {
// 		s.CO2 = 0
// 	}

// 	return s
// }

// func (s Spaceship) Eat(f int) (Spaceship, bool) {
// 	s.Food = s.Food - f
// 	if s.Food > 0 {
// 		return s, true
// 	} else {
// 		s.Food = 0
// 		return s, false
// 	}
// }

// func (s Spaceship) Drink(f int) (Spaceship, bool) {
// 	s.Water = s.Water - f
// 	if s.Water > 0 {
// 		return s, true
// 	} else {
// 		s.Water = 0
// 		return s, false
// 	}
// }
