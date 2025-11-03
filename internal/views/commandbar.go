package views

import (
	"dctui/internal/theme"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CommandBarView struct {
	Primitive *tview.InputField
}

func NewCommandBarView() *CommandBarView {
	labelStyle := tcell.StyleDefault
	labelStyle.Background(tcell.ColorBlue)
	input := tview.NewInputField().SetLabel("> ").SetFieldWidth(0)
	input.SetBorder(true).SetBorderColor(tcell.ColorLightBlue)
	input.SetBackgroundColor(theme.BgColor)
	input.SetFieldBackgroundColor(theme.BgColor)
	return &CommandBarView{Primitive: input}
}
