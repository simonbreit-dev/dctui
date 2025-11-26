package ui

import (
	"dctui/internal/theme"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainViewMode int

const (
	ModeList MainViewMode = iota
	ModeDetail
)

type MainView struct {
	app         *tview.Application
	table       *tview.Table
	mode        MainViewMode
	OnSelectRow func(row int, col int)
}

func NewMainView(app *tview.Application) *MainView {
	table := tview.NewTable()
	table.SetSelectable(true, false).
		SetBorder(true).
		SetBackgroundColor(theme.BgColor)

	view := &MainView{
		app:   app,
		table: table,
		mode:  ModeList,
	}

	table.SetSelectedFunc(func(row, col int) {
		if view.OnSelectRow != nil {
			view.OnSelectRow(row, col)
		}
	})

	return view
}

func (v *MainView) GetPrimitive() tview.Primitive {
	return v.table
}

func (v *MainView) RenderWithData(containers []container.Summary) {
	v.mode = ModeList
	v.table.Clear()

	// Header
	v.table.SetCell(0, 0, tview.NewTableCell("ID").SetAttributes(tcell.AttrBold))
	v.table.SetCell(0, 1, tview.NewTableCell("Name").SetAttributes(tcell.AttrBold))
	v.table.SetCell(0, 2, tview.NewTableCell("Image").SetAttributes(tcell.AttrBold))
	v.table.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrBold))

	for i, c := range containers {
		// IDs kürzen
		id := c.ID
		if len(id) > 12 {
			id = id[:12]
		}

		// Namen kombinieren, da c.Names ein Array ist
		name := strings.Join(c.Names, ", ")

		v.table.SetCell(i+1, 0, tview.NewTableCell(id))
		v.table.SetCell(i+1, 1, tview.NewTableCell(name))
		v.table.SetCell(i+1, 2, tview.NewTableCell(c.Image))
		v.table.SetCell(i+1, 3, tview.NewTableCell(c.Status))
	}

	// SelectedFunc für Zeilen
	v.table.SetSelectable(true, false)
	if v.OnSelectRow != nil {
		v.table.SetSelectedFunc(v.OnSelectRow)
	}
}

func (v *MainView) RenderDetail(info string) {
	v.mode = ModeDetail
	v.table.Clear()
	v.table.SetCell(0, 0, tview.NewTableCell(info))
}

func (v *MainView) IsDetailMode() bool {
	return v.mode == ModeDetail
}

// ContainerData ist eine einfache Struktur, die Controller an View liefert
type ContainerData struct {
	ID     string
	Image  string
	Status string
}
