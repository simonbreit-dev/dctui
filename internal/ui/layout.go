package ui

import (
	"github.com/rivo/tview"
)

type LayoutView struct {
	app        *tview.Application
	grid       *tview.Grid
	header     *HeaderView
	mainArea   *tview.Flex // Hier werden die Seiten geladen
	commandBar *CommandBarView
}

func NewLayoutView(header *HeaderView, commandBar *CommandBarView, app *tview.Application) *LayoutView {
	mainArea := tview.NewFlex()
	mainArea.SetDirection(tview.FlexRow)

	grid := tview.NewGrid().
		SetRows(3, 1, 0). // Header, MainArea, CommandBar
		SetColumns(0).
		AddItem(header.GetPrimitive(), 0, 0, 1, 1, 0, 0, false).
		AddItem(commandBar.GetPrimitive(), 1, 0, 1, 1, 0, 0, false).
		AddItem(mainArea, 2, 0, 1, 1, 0, 0, true)

	return &LayoutView{
		grid:       grid,
		header:     header,
		mainArea:   mainArea,
		commandBar: commandBar,
		app:        app,
	}
}

func (l *LayoutView) GetPrimitive() tview.Primitive {
	return l.grid
}

func (l *LayoutView) SetMainPage(p tview.Primitive) {
	l.mainArea.Clear()
	l.mainArea.AddItem(p, 0, 1, true)
}

func (l *LayoutView) ExpandCommandBar() {
	l.grid.SetRows(3, 3, 0) // CommandBar größer
	l.app.SetFocus(l.commandBar.GetPrimitive())
}

func (l *LayoutView) CollapseCommandBar() {
	l.app.SetFocus(l.mainArea)
	l.grid.SetRows(3, 1, 0) // CommandBar wieder klein
}
