package main

import (
	"log"
	"strconv"
	"time"

	en "github.com/dbenoot/moonbase/engine"
	"github.com/marcusolsson/tui-go"
	"github.com/kutase/go-gameloop"
)

var engineOutput string
var turn int
var base en.Moonbase

func main() {



	base = en.Start()

	sbcontent := tui.NewLabel("SIDEBAR")

	sbcontent.SetText(base.Name + "\n" + base.Government + "\n" + base.Sponsor + "\n\n" + strconv.Itoa(base.Money) + "\n" + strconv.Itoa(base.Health) + "\n" + strconv.Itoa(base.Lifesupport) + "\n\n" + strconv.Itoa(base.Water) + "\n" + strconv.Itoa(base.Food) + "\n" + strconv.Itoa(base.Fuel))

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

	gl := gameLoop.New(10, func(delta float64) {
		input.OnSubmit(func(e *tui.Entry) {

			engineOutput, turn, base = en.Input(e.Text())

			main.Append(tui.NewHBox(
				tui.NewLabel("turn "+strconv.Itoa(turn)),
				tui.NewPadder(1, 0, tui.NewLabel(" - ")),
				tui.NewLabel(time.Now().Format("15:04")),
				tui.NewPadder(1, 0, tui.NewLabel(" >")),
				tui.NewLabel(engineOutput),
				tui.NewSpacer(),
			))
			input.SetText("")

		})

		sbcontent.SetText(base.Name + "\n" + base.Government + "\n" + base.Sponsor + "\n\n" + strconv.Itoa(base.Money) + "\n" + strconv.Itoa(base.Health) + "\n" + strconv.Itoa(base.Lifesupport) + "\n\n" + strconv.Itoa(base.Water) + "\n" + strconv.Itoa(base.Food) + "\n" + strconv.Itoa(base.Fuel) + "\n\n\n" + "tick:" +  strconv.FormatFloat(delta, 'f', 6, 64	))

		})

	gl.Start()



	root := tui.NewHBox(sidebar, command)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
