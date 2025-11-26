package ui

import (
	service "dctui/internal/services"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIController struct {
	app        *tview.Application
	layout     *LayoutView
	pages      map[string]Page
	activePage Page
	history    []string
	docker     *service.DockerService
}

func NewUIController(app *tview.Application, layout *LayoutView, docker *service.DockerService) *UIController {
	c := &UIController{
		app:    app,
		layout: layout,
		pages:  make(map[string]Page),
		docker: docker,
	}
	c.setupInputCapture() // Globale Shortcuts aktivieren
	return c
}

func (c *UIController) AddPage(name string, p Page) {
	c.pages[name] = p
	switch casted := p.(type) {
	case *ProjectOverview:
		casted.OnSelectRow = func(row int, col int, projectName string) {
			containers := c.docker.FetchContainersForProject(projectName)
			c.SwitchToPage("projectView", containers)
		}
	case *ProjectView:

	}
}

func (c *UIController) SwitchToPage(name string, data any) {
	if p, ok := c.pages[name]; ok {
		// Aktuelle Seite speichern, nur wenn es eine echte Seite ist
		if c.activePage != nil && c.activePage != p {
			c.history = append(c.history, c.getActivePageName())
		}

		if c.activePage != nil {
			c.activePage.OnBlur()
		}

		c.activePage = p
		if data != nil {
			p.RenderWithData(data)
		}
		c.activePage.OnFocus()
		c.layout.SetMainPage(p.GetPrimitive())
		c.app.SetFocus(p.GetPrimitive())
	}
}

func (c *UIController) setupInputCapture() {
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			if len(c.history) > 0 {
				// Letzte Seite aus dem Stack holen
				prev := c.history[len(c.history)-1]
				c.history = c.history[:len(c.history)-1]
				c.SwitchToPage(prev, nil)
				return nil
			}
		}

		// Globale Shortcuts
		switch event.Rune() {
		case ':':
			c.layout.ExpandCommandBar()
		}
		return event
	})
}

func (c *UIController) CloseCommandBar() {
	c.layout.CollapseCommandBar()
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

func (c *UIController) getActivePageName() string {
	for name, page := range c.pages {
		if page == c.activePage {
			return name
		}
	}
	return ""
}
