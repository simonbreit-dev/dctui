package ui

import (
	"dctui/internal/theme"
	"github.com/rivo/tview"
)

type LayoutView struct {
	app        *tview.Application
	grid       *tview.Grid
	header     *HeaderView
	commandBar *CommandBarView
	mainView   *MainView
}

func NewLayoutView(app *tview.Application, header *HeaderView, commandBar *CommandBarView, mainView *MainView) *LayoutView {
	grid := tview.NewGrid().
		SetRows(5, 1, 0).
		SetColumns(0).
		SetBorders(false)
	grid.AddItem(header.GetPrimitive(), 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(commandBar.GetPrimitive(), 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(mainView.GetPrimitive(), 2, 0, 1, 1, 0, 0, true)
	grid.SetBackgroundColor(theme.BgColor)
	return &LayoutView{app: app, grid: grid, header: header, commandBar: commandBar, mainView: mainView}
}

func (v *LayoutView) ExpandCommandBar() {
	v.grid.SetRows(5, 3, 0)
}

func (v *LayoutView) CollapseCommandBar() {
	v.grid.SetRows(5, 1, 0)
}

func (v *LayoutView) SetCommandBar(cb *CommandBarView) {
	v.commandBar = cb
}

func (v *LayoutView) GetPrimitive() tview.Primitive {
	return v.grid
}
