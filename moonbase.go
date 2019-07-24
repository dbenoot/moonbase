package main

import (
	"log"
	"time"

	en "github.com/dbenoot/moonbase/engine"
	"github.com/marcusolsson/tui-go"
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

	astrocontent := tui.NewLabel(en.GetAstroNames())
	astroBox := tui.NewHBox(astrocontent)
	astroBox.SetBorder(true)
	astroBox.SetSizePolicy(tui.Minimum, tui.Minimum)

	locationcontent := tui.NewLabel(en.GetLocations())
	locationBox := tui.NewHBox(locationcontent)
	locationBox.SetBorder(true)
	locationBox.SetSizePolicy(tui.Minimum, tui.Minimum)

	infoBox := tui.NewVBox(astroBox, locationBox)
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
				locationcontent.SetText(en.GetLocations())
				astrocontent.SetText(en.GetAstroNames())
			})

		}
	}()

	// Send commands to the engine update the sidebar.
	// The sidebar is updated directly, should there be any changes in the bases' stats

	go input.OnSubmit(func(e *tui.Entry) {

		en.Input(e.Text())
		input.SetText("")
		sbcontent.SetText(en.GetSideBarInfo())

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
				tui.NewLabel(mainOutput),
				tui.NewSpacer(),
			))

		}
	}()

	// Now that the interface is set up (especially the channel listeners), start the engine

	en.Start()

	// Check errors

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
