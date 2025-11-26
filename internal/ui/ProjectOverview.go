package ui

import (
	"dctui/internal/models"
	"dctui/internal/theme"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectOverview struct {
	app         *tview.Application
	table       *tview.Table
	OnSelectRow func(row int, col int, value string)
}

func (v *ProjectOverview) OnFocus() {
	return
}

func (v *ProjectOverview) OnBlur() {
	return
}

func NewProjectOverview(app *tview.Application) *ProjectOverview {
	table := tview.NewTable()
	table.SetSelectable(true, false).
		SetBorder(true).
		SetBackgroundColor(theme.BgColor)

	view := &ProjectOverview{
		app:   app,
		table: table,
	}

	table.SetSelectedFunc(func(row, col int) {
		if row == 0 {
			return // Header ignorieren
		}
		// Erstes Cell (Spalte 0)
		cell := table.GetCell(row, 0)
		projectName := cell.Text
		table.SetSelectable(true, false)
		if view.OnSelectRow != nil {
			view.OnSelectRow(row, col, projectName)
		}
		table.SetSelectedFunc(func(row, col int) {
			// Header auslassen
			if row == 0 {
				return
			}
			cell := table.GetCell(row, 0)
			value := cell.Text

			if view.OnSelectRow != nil {
				view.OnSelectRow(row, col, value)
			}
		})
	})

	return view
}

func (v *ProjectOverview) GetPrimitive() tview.Primitive {
	return v.table
}

func (v *ProjectOverview) RenderWithData(data any) {
	projects, err := data.([]models.Project)
	if !err {
		return
	}
	v.table.Clear()

	// Header
	v.table.SetCell(0, 0, tview.NewTableCell("Name").SetAttributes(tcell.AttrBold).SetSelectable(false)) // 👈 wichtig!
	v.table.SetCell(0, 1, tview.NewTableCell("Containers").SetAttributes(tcell.AttrBold).SetSelectable(false))
	v.table.SetCell(0, 2, tview.NewTableCell("File").SetAttributes(tcell.AttrBold).SetSelectable(false))
	v.table.SetCell(0, 3, tview.NewTableCell("Working Directory").SetAttributes(tcell.AttrBold).SetSelectable(false))

	for i, p := range projects {
		v.table.SetCell(i+1, 0, tview.NewTableCell(p.Name))
		count := strconv.Itoa(p.ContainersCount)
		v.table.SetCell(i+1, 1, tview.NewTableCell(count))
		v.table.SetCell(i+1, 2, tview.NewTableCell(p.ConfigFile))
		v.table.SetCell(i+1, 3, tview.NewTableCell(p.WorkingDir))
	}

	// SelectedFunc für Zeilen

}
