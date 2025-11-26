package ui

import (
	service "dctui/internal/services"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIController struct {
	app        *tview.Application
	docker     *service.DockerService
	mainView   *MainView
	header     *HeaderView
	commandBar *CommandBarView
	layout     *LayoutView
}

func NewUIController(app *tview.Application, docker *service.DockerService,
	mainView *MainView, header *HeaderView, commandBar *CommandBarView, layout *LayoutView) *UIController {

	c := &UIController{
		app:        app,
		docker:     docker,
		mainView:   mainView,
		header:     header,
		commandBar: commandBar,
		layout:     layout,
	}

	// MainView Auswahl
	mainView.OnSelectRow = func(row int, col int) {
		if row == 0 { // Header-Zeile ignorieren
			return
		}
		containerID := mainView.table.GetCell(row, 0).Text
		info := c.getContainerInfo(containerID)
		c.mainView.RenderDetail(info)
		c.app.SetFocus(c.mainView.GetPrimitive())
	}

	c.setupInputCapture()

	go c.RefreshMainView()

	return c
}

func (c *UIController) GetLayout() tview.Primitive {
	return c.layout.GetPrimitive()
}

func (c *UIController) SetCommandBar(cb *CommandBarView) {
	c.commandBar = cb
}

func (c *UIController) OpenCommandBar() {
	c.layout.ExpandCommandBar()
	c.commandBar.Clear()
	c.app.SetFocus(c.commandBar.GetPrimitive())
}

func (c *UIController) CloseCommandBar() {
	c.layout.CollapseCommandBar()
	c.app.SetFocus(c.mainView.GetPrimitive())
}

func (c *UIController) ExecuteCommand(cmd string) {
	cmd = strings.ToLower(cmd)
	cmd = strings.TrimSpace(cmd)
	cmd = strings.Trim(cmd, ":")
	if cmd == "quit" || cmd == "exit" || cmd == "q" || cmd == "q!" {
		c.app.Stop()
	}

	c.CloseCommandBar()
}

func (c *UIController) setupInputCapture() {
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// ESC Detailansicht schließen
		if c.mainView.IsDetailMode() && event.Key() == tcell.KeyEsc {
			c.mainView.RenderWithData(c.docker.FetchContainers())
			c.app.SetFocus(c.mainView.GetPrimitive())
			return nil
		}

		// Globale Shortcuts
		switch event.Rune() {
		case ':':
			c.OpenCommandBar()
			c.app.SetFocus(c.commandBar.GetPrimitive())
		case 'r':
			c.RefreshMainView()
		case '1':
			c.app.SetFocus(c.mainView.GetPrimitive())
		case '3':
			c.app.SetFocus(c.header.GetPrimitive())
		}
		return event
	})
}

func (c *UIController) RefreshMainView() {
	// Service holt die Rohdaten
	raw := c.docker.FetchContainers()

	// Mapping auf View-kompatible Struktur
	data := make([]ContainerData, len(raw))
	for i, ctn := range raw {
		data[i] = ContainerData{
			ID:     ctn.ID[:12],
			Image:  ctn.Image,
			Status: ctn.Status,
		}
	}

	// View aktualisieren im UI-Goroutine
	c.app.QueueUpdateDraw(func() {
		c.mainView.RenderWithData(raw)
	})
}

// Container-Details von DockerService holen
func (c *UIController) getContainerInfo(containerID string) string {
	rawContainers := c.docker.FetchContainers()
	for _, ctn := range rawContainers {
		if ctn.ID[:12] == containerID {
			return fmt.Sprintf("ID: %s\nImage: %s\nStatus: %s", ctn.ID, ctn.Image, ctn.Status)
		}
	}
	return "Container nicht gefunden"
}
