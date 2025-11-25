package ui

import (
	"dctui/internal/theme"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CommandBarView struct {
	app        *tview.Application
	input      *tview.InputField
	controller *UIController
}

func NewCommandBarView(app *tview.Application) *CommandBarView {
	input := tview.NewInputField().SetLabel("> ")
	input.SetBorder(true).SetBorderColor(tcell.ColorLightBlue)
	input.SetBackgroundColor(theme.BgColor)
	view := &CommandBarView{app: app, input: input}
	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			if view.controller != nil {
				view.controller.ExecuteCommand(input.GetText())
			}
		} else if key == tcell.KeyEsc {
			if view.controller != nil {
				view.controller.CloseCommandBar()
			}
		}
	})
	return view
}

func (v *CommandBarView) SetController(c *UIController) {
	v.controller = c
}

func (v *CommandBarView) Clear() {
	v.input.SetText("")
}

func (v *CommandBarView) GetPrimitive() tview.Primitive {
	return v.input
}
