package main

import (
	"log"
	"time"
	"sort"

	en "github.com/dbenoot/moonbase/engine"
	"github.com/marcusolsson/tui-go"
	"github.com/marcusolsson/tui-go/wordwrap"
)

var engineOutput string

func main() {

	// set up the interface

	sbcontent := tui.NewLabel(en.GetSideBarInfo())

	sidebar := tui.NewVBox(
		sbcontent,
		tui.NewSpacer(),
	)
	sidebar.SetBorder(true)

	main := tui.NewVBox()

	mainScroll := tui.NewScrollArea(main)
	mainScroll.SetAutoscrollToBottom(true)

	mainBox := tui.NewVBox(mainScroll)
	mainBox.SetBorder(true)

	mapcontent := tui.NewLabel("")
	mapBox := tui.NewHBox(mapcontent)
	mapBox.SetBorder(true)
	mapBox.SetSizePolicy(tui.Expanding, tui.Minimum)
	mapBox.SetTitle("Map")

	playercontent := tui.NewLabel(en.GetPlayerStats())
	playerBox := tui.NewHBox(playercontent)
	playerBox.SetTitle("Stats")
	playerBox.SetBorder(true)
	playerBox.SetSizePolicy(tui.Minimum, tui.Minimum)

	infoBox := tui.NewVBox(mapBox, playerBox)
	infoBox.SetSizePolicy(tui.Minimum, tui.Expanding)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	command := tui.NewVBox(mainBox, inputBox)
	command.SetSizePolicy(tui.Expanding, tui.Expanding)

	root := tui.NewHBox(sidebar, command, infoBox)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}	

	// Start a poller for interface updates. Polling starts from the interface now. Update to channel as in themain window?

	go func() {
		for range time.Tick(time.Second * 1) {
			ui.Update(func() {
				sbcontent.SetText(en.GetSideBarInfo())
				playercontent.SetText(en.GetPlayerStats())	 // moved here as the map doesn't exist at start and then minMaxIntSlice crashes, update update to channel from engine!			
			})

		}
	}()

	// Send commands to the engine update the sidebar.
	// The sidebar is updated directly, should there be any changes in the bases' stats

	go input.OnSubmit(func(e *tui.Entry) {

		en.Input(e.Text())
		input.SetText("")
		sbcontent.SetText(en.GetSideBarInfo())
		mapcontent.SetText(drawMap())

	})

	// Listener set up for the quit signal

	go func() {
		quit, _ := <-en.Quit

		if quit == true {
			ui.Quit()
		}
	}()

	// Listener for the engine output

	go func() {
		for {
			mainOutput, _ := <-en.Output

			main.Append(tui.NewHBox(
				tui.NewPadder(1, 0, tui.NewLabel(">")),
				tui.NewLabel(wordwrap.WrapString(mainOutput, mainBox.Size().X-10)),
				tui.NewSpacer(),
			))

		}
	}()

	// Now that the interface is set up (especially the channel listeners), ask the user to start the engine

	main.Append(tui.NewHBox(
		tui.NewPadder(1, 0, tui.NewLabel(">")),
		tui.NewLabel("Type 'start' to start the game."),
		tui.NewSpacer(),
	))
	
	// Check errors

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func drawMap() string {
	lm := en.GetMap()
	pl := en.GetProtLoc()

	var xxx []int
	var yyy []int

	var output string

	for k, _ := range lm {
		xxx = append(xxx, k.X)
		yyy = append(yyy, k.Y)
	}
	minx, maxx := minMaxIntSlice(xxx)
	miny, maxy := minMaxIntSlice(yyy)

	for y := maxy+1; y > miny-1; y-- {
		for x := minx-1; x < maxx+1; x++ {
			_, ok := lm[en.Coordinates{x, y}]
			if ok && pl.X == x && pl.Y == y {
				output = output + "@"
			} else if ok {
				output = output + "X"
			} else {
				output = output + " "
			}
		}
		output = output + "\n"
	}

	return output
}

func minMaxIntSlice (v []int) (int, int) {
	sort.Ints(v)
	return v[0], v[len(v)-1]
}