package main

import (
	"dctui/internal/services"
	"dctui/internal/ui"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	dockerService := services.NewDockerService()

	header := ui.NewHeaderView()
	mainView := ui.NewMainView(app, dockerService)
	commandBar := ui.NewCommandBarView(app)
	layout := ui.NewLayoutView(app, header, commandBar, mainView)

	controller := ui.NewUIController(app, layout, header, mainView, commandBar)
	commandBar.SetController(controller)

	if err := app.SetRoot(layout.GetPrimitive(), true).Run(); err != nil {
		panic(err)
	}
}
