package ui

import (
	"dctui/internal/theme"

	"github.com/rivo/tview"
)

type HeaderView struct {
	primitive *tview.TextView
}

func NewHeaderView() *HeaderView {
	tv := tview.NewTextView()
	tv.SetText("VM:\nRAM:\nCPU:\n ")
	tv.SetBackgroundColor(theme.BgColor)
	return &HeaderView{primitive: tv}
}

func (v *HeaderView) Render() {}

func (v *HeaderView) GetPrimitive() tview.Primitive { return v.primitive }
