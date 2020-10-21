package engine

import (
	"math/rand"
	"time"
)

// Astronaut process memory function

func (a *Astronaut) processMemory() {
	// Then the STM is being forgot at a rate of 2

	a.processStm()

	// Then the LTM is being forgot at a rate of 1
	// TODO: and the oldest memories just stay at zero, they don't get deleted.

	a.processLtm()

	// every 10 minutes copy the AM to the STM
	// TODO: %600 is an hour, not 10 minutes!

	if t%600 == 0 {
		a.activeToStm()
	}

	// every hour copy the STM to the LTM
	// TODO: %3600 is every day, not every hour!

	if t%3600 == 0 {
		a.imprint()
	}
}

// Basic

func (m *Memory) decreasepersistence(a int) {
	m.persistence = m.persistence - a
}

// Processing

func (a *Astronaut) processLtm() {

	for i := range a.Ltm {
		if a.Ltm[i].persistence > 1 {
			a.Ltm[i].decreasepersistence(1)
		}
	}

}

func (a *Astronaut) processStm() {

	for i := 0; i < len(a.Stm); i++ {
		a.Stm[i].decreasepersistence(1)

		if a.Stm[i].memory == a.Activemem.memory {
			a.Stm[i].persistence = a.Stm[i].persistence + 3 // TODO make variables list at the top to make balancing easier
		}
		if a.Stm[i].persistence == 0 {
			a.Stm, i = removeMem(a.Stm, i) // functions lowers i by 1 otherwise we have an out of bounds panic
		}
	}

}

// Actions

func removeMem(s []Memory, i int) ([]Memory, int) {
	return append(s[:i], s[i+1:]...), i - 1
}

func (a *Astronaut) addNewActiveMem(mem string, desc string, q int) {
	a.Activemem = Memory{mem, desc, q, 30}
	a.activeToStm()
}

func (a *Astronaut) addActiveMem(mem Memory) {
	a.Activemem = mem
	a.activeToStm()
}

func (a *Astronaut) reminisce() {
	// select a memory from long-term memory and send it to the active memory
	// TODO: this should work in such a way that e.g. top 10% of the ltm has 50% chance of returning, while the bottom 50% only 10% for example

	rand.Seed(time.Now().Unix())
	memory := a.Ltm[rand.Intn(len(a.Ltm))]

	a.addActiveMem(memory)
	Output <- a.Name + " is remembering an old memory: " + memory.memory
}

func (a *Astronaut) activeToStm() {
	var knownmem = false
	for i := range a.Stm {
		if a.Activemem.memory == a.Stm[i].memory {
			knownmem = true
		}
	}

	if knownmem == true {
		for i := range a.Stm {
			if a.Activemem.memory == a.Stm[i].memory {
				a.Stm[i].persistence = a.Stm[i].persistence + 3
			}
		}
	} else {
		a.Stm = append(a.Stm, a.Activemem)
	}

}

func (a *Astronaut) imprint() {
	// All short term items are copied to the long term memory with their countdown time.
	// Countdown time is added to the long term memory
	// Countdown time for remembered items must thus be equal to passed time
	// By doing so, very memorable items will remain in the ltm
	// Q: copy every cycle? No, because timer should have time to run out. Every ingame hour or so

	for i, mem := range a.Stm {
		k, found := findMem(a.Ltm, mem)
		if !found {
			a.Ltm = append(a.Ltm, a.Stm[i])
		} else {
			a.Ltm[k].persistence = a.Ltm[k].persistence + 3
		}
	}
}

func findMem(s []Memory, v Memory) (int, bool) {
	for i, mem := range s {
		if mem == v {
			return i, true
		}
	}
	return -1, false
}
