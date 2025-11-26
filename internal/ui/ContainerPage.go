package ui

import (
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ContainerListPage struct {
	table    *tview.Table
	OnSelect func(containerID string)
}

func NewContainerListPage(containers []container.Summary, onSelect func(string)) *ContainerListPage {
	t := tview.NewTable().SetSelectable(true, false)
	page := &ContainerListPage{
		table:    t,
		OnSelect: onSelect,
	}

	// Header
	t.SetCell(0, 0, tview.NewTableCell("ID").SetAttributes(tcell.AttrBold))
	t.SetCell(0, 1, tview.NewTableCell("Name").SetAttributes(tcell.AttrBold))
	t.SetCell(0, 2, tview.NewTableCell("Image").SetAttributes(tcell.AttrBold))
	t.SetCell(0, 3, tview.NewTableCell("Status").SetAttributes(tcell.AttrBold))

	for i, c := range containers {
		id := c.ID[:12]
		name := strings.Join(c.Names, ", ")
		t.SetCell(i+1, 0, tview.NewTableCell(id))
		t.SetCell(i+1, 1, tview.NewTableCell(name))
		t.SetCell(i+1, 2, tview.NewTableCell(c.Image))
		t.SetCell(i+1, 3, tview.NewTableCell(c.Status))
	}

	t.SetSelectedFunc(func(row, column int) {
		if row == 0 {
			return
		}
		containerID := t.GetCell(row, 0).Text
		if page.OnSelect != nil {
			page.OnSelect(containerID)
		}
	})

	return page
}

func (p *ContainerListPage) GetPrimitive() tview.Primitive { return p.table }
func (p *ContainerListPage) OnFocus()                      {}
func (p *ContainerListPage) OnBlur()                       {}
