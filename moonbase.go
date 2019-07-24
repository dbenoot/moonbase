package main

import (
	"log"
	"time"

	en "github.com/dbenoot/moonbase/engine"
	"github.com/marcusolsson/tui-go"
)

var engineOutput string

func main() {

	en.Start()

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

	// ui.SetKeybinding("Esc", func() { ui.Quit() })

	go func() {
		for range time.Tick(time.Second * 1) {
			ui.Update(func() {
				sbcontent.SetText(en.GetSideBarInfo())
				locationcontent.SetText(en.GetLocations())
				astrocontent.SetText(en.GetAstroNames())
			})
		}
	}()

	go input.OnSubmit(func(e *tui.Entry) {

		en.Input(e.Text())
		input.SetText("")
		sbcontent.SetText(en.GetSideBarInfo())

	})

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

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
