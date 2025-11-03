package views

import (
	"dctui/internal/theme"

	"github.com/rivo/tview"
)

type HeaderView struct {
	Primitive *tview.TextView
}

func NewHeaderView() *HeaderView {
	tv := tview.NewTextView()
	tv.SetText("VM:\nRAM:\nCPU:\n ")
	tv.SetBackgroundColor(theme.BgColor)
	return &HeaderView{Primitive: tv}
}
