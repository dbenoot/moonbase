package engine

func look() {
	for _, v := range locations {
		if v.Name == allAstronauts[0].Location {
			Output <- v.Description
		}
	}
}
