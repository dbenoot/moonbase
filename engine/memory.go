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

	a.processLtm()

	// every 10 minutes copy the AM to the STM

	if t%600 == 0 {
		a.activeToStm()
	}

	// every hour copy the STM to the LTM

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

	for i := range a.Stm {
		a.Stm[i].decreasepersistence(1)

		if a.Stm[i].memory == a.Activemem.memory {
			a.Stm[i].persistence = a.Stm[i].persistence + 3
		}
		/* if a.Ltm[i].persistence == 0 {
			a.Ltm[len(a.Ltm)-1], a.Ltm[i] = a.Ltm[i], a.Ltm[len(a.Ltm)-1]
			a.Ltm = a.Ltm[:len(a.Ltm)-1]
		} */
	}

}

// Actions

func (a *Astronaut) addActiveMem(mem string, desc string, q int) {
	a.Activemem = Memory{mem, desc, q, 30}
}

func (a *Astronaut) reminisce() {
	// select a memory from long-term memory and send it to the active memory
	// this should work in such a way that e.g. top 10% of the ltm has 50% chance of returning, while the bottom 50% only 10% for example

	rand.Seed(time.Now().Unix())
	memory := a.Ltm[rand.Intn(len(a.Ltm))]

	a.Activemem = memory
}

func (a *Astronaut) activeToStm() {
	var knownmem bool = false
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

}
