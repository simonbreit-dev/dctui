package views

import (
	"dctui/internal/state"
	"dctui/internal/theme"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainView struct {
	Primitive *tview.Table
}

func NewMainView() *MainView {
	table := tview.NewTable()
	table.SetBorder(true)   // behalte den Rahmen
	table.SetBorders(false) // entfernt interne Linien zwischen Zellen
	table.SetBorderColor(tcell.ColorRed)
	table.SetBackgroundColor(theme.BgColor)
	return &MainView{Primitive: table}
}

func (m *MainView) Render(state *state.AppState) {
	m.Primitive.Clear()
	for i, item := range state.Items {
		m.Primitive.SetCell(i, 0, tview.NewTableCell(item))
	}
}
