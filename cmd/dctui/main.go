package main

import (
	"dctui/internal/appinstance"
	"dctui/internal/controller"
	"dctui/internal/state"
	"dctui/internal/theme"
	"dctui/internal/views"

	"github.com/rivo/tview"
)

func main() {
	appState := state.NewAppState()
	app := tview.NewApplication()
	appinstance.App = app
	// Views
	header := views.NewHeaderView()
	mainView := views.NewMainView()
	commandBar := views.NewCommandBarView()

	// Body Flex: Sidebar + MainContent
	bodyFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(mainView.Primitive, 0, 4, false)

	// Root Grid: Header, Body, CommandBar
	rootGrid := tview.NewGrid().
		SetRows(5, 3, 0). // Header, Commandbar, Main (flexibel)
		SetColumns(0).    // volle Breite
		SetBorders(false).
		AddItem(header.Primitive, 0, 0, 1, 1, 0, 0, false).
		AddItem(commandBar.Primitive, 1, 0, 1, 1, 0, 0, false).
		AddItem(bodyFlex, 2, 0, 1, 1, 0, 0, false)
	rootGrid.SetBackgroundColor(theme.BgColor)

	// Controller
	controller.NewController(app, appState, mainView, commandBar, rootGrid)

	// Start App
	if err := app.SetRoot(rootGrid, true).SetFocus(mainView.Primitive).Run(); err != nil {
		panic(err)
	}
}
