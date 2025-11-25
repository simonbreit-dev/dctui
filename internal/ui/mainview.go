package ui

import (
	"dctui/internal/services"
	"dctui/internal/theme"

	"github.com/rivo/tview"
)

type MainView struct {
	app    *tview.Application
	docker *services.DockerService
	table  *tview.Table
}

func NewMainView(app *tview.Application, docker *services.DockerService) *MainView {
	table := tview.NewTable()
	table.SetBorder(true).SetBackgroundColor(theme.BgColor)
	view := &MainView{app: app, docker: docker, table: table}
	view.Render()
	return view
}

func (v *MainView) Render() {
	v.table.Clear()
}

func (v *MainView) GetPrimitive() tview.Primitive {
	return v.table
}
