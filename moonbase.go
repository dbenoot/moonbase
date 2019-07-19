package main

import (
	"log"
	"strconv"
	"time"

	en "github.com/dbenoot/moonbase/engine"
	"github.com/marcusolsson/tui-go"
)

var engineOutput string
var turn int

func main() {

	en.Start()

	sbcontent := tui.NewLabel("SIDEBAR")

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

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	command := tui.NewVBox(mainBox, inputBox)
	command.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {

		engineOutput = en.Input(e.Text())

		main.Append(tui.NewHBox(
			tui.NewLabel("turn "+strconv.Itoa(turn)),
			tui.NewPadder(1, 0, tui.NewLabel(" - ")),
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(" >")),
			tui.NewLabel(engineOutput),
			tui.NewSpacer(),
		))
		input.SetText("")
		sbcontent.SetText(en.GetSideBarInfo())
	})

	root := tui.NewHBox(sidebar, command)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	ui.SetKeybinding("Space", func() { en.PauseUnPause() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
