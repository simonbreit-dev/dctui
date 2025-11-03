package controller

import (
	"dctui/internal/state"
	"dctui/internal/views"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Controller struct {
	app        *tview.Application
	state      *state.AppState
	mainView   *views.MainView
	commandBar *views.CommandBarView
	rootGrid   *tview.Grid
}

func NewController(app *tview.Application, state *state.AppState, mainView *views.MainView, commandBar *views.CommandBarView, rootGrid *tview.Grid) *Controller {
	c := &Controller{
		app:        app,
		state:      state,
		mainView:   mainView,
		commandBar: commandBar,
		rootGrid:   rootGrid,
	}

	c.setupCommandBar()
	c.setupGlobalShortcuts()
	c.mainView.Render(c.state)

	return c
}

func (c *Controller) setupCommandBar() {
	c.closeCommandBar()
	c.commandBar.Primitive.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			executeCommand(c.commandBar.Primitive.GetText())
			c.closeCommandBar()
		case tcell.KeyEsc:
			c.closeCommandBar()
		}
	})
}

func (c *Controller) setupGlobalShortcuts() {
	c.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'j':
				c.state.SelectedIdx++
			case 'k':
				c.state.SelectedIdx--
			case ':':
				c.openCommandBar()
				return nil
			}
		}
		return event
	})
}

func (c *Controller) openCommandBar() {
	c.rootGrid.SetRows(5, 3, 0)
	c.state.CommandMode = true
	c.app.SetFocus(c.commandBar.Primitive)
	c.commandBar.Primitive.SetText("")
}

func (c *Controller) closeCommandBar() {
	c.rootGrid.SetRows(5, 1, 0)
	c.state.CommandMode = false
	c.app.SetFocus(c.mainView.Primitive)
}
