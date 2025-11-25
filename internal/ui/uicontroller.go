package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIController struct {
	app        *tview.Application
	layout     *LayoutView
	header     *HeaderView
	mainView   *MainView
	commandBar *CommandBarView
}

func NewUIController(app *tview.Application, layout *LayoutView, header *HeaderView, mainView *MainView, commandBar *CommandBarView) *UIController {
	c := &UIController{app: app, layout: layout, header: header, mainView: mainView, commandBar: commandBar}
	c.setupGlobalShortcuts()
	return c
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
	// Handle command, e.g., call DockerService
	// Placeholder for future logic
	_ = cmd // to avoid unused variable warning
	c.CloseCommandBar()
}

func (c *UIController) setupGlobalShortcuts() {
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case ':':
			c.OpenCommandBar()
		case '1':
			c.app.SetFocus(c.mainView.GetPrimitive())
		case '3':
			c.app.SetFocus(c.header.GetPrimitive())
		}
		return event
	})
}
