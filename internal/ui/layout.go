package ui

import (
	"dctui/internal/theme"

	"github.com/rivo/tview"
)

type LayoutView struct {
	grid       *tview.Grid
	header     tview.Primitive
	mainView   tview.Primitive
	commandBar tview.Primitive
}

func NewLayoutView(header, mainView, commandBar tview.Primitive) *LayoutView {
	grid := tview.NewGrid().
		SetRows(3, 1, 0). // Header, Main, CommandBar
		SetColumns(0).
		AddItem(header, 0, 0, 1, 1, 0, 0, false).
		AddItem(commandBar, 1, 0, 1, 1, 0, 0, false).
		AddItem(mainView, 2, 0, 1, 1, 0, 0, true)

	grid.SetBackgroundColor(theme.BgColor)

	return &LayoutView{
		grid:       grid,
		header:     header,
		mainView:   mainView,
		commandBar: commandBar,
	}
}

func (l *LayoutView) GetPrimitive() tview.Primitive {
	return l.grid
}

// CommandBar expand/collapse falls gewünscht
func (l *LayoutView) ExpandCommandBar() {
	l.grid.SetRows(3, 3, 0) // CommandBar größer
}

func (l *LayoutView) CollapseCommandBar() {
	l.grid.SetRows(3, 1, 0) // CommandBar wieder klein
}
