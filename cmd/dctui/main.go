package main

import (
	service "dctui/internal/services"
	"dctui/internal/ui"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	docker := service.NewDockerService()

	header := ui.NewHeaderView()
	mainView := ui.NewMainView(app)
	commandBar := ui.NewCommandBarView(app)
	layout := ui.NewLayoutView(header.GetPrimitive(), mainView.GetPrimitive(), commandBar.GetPrimitive())
	controller := ui.NewUIController(app, docker, mainView, header, commandBar, layout)

	commandBar.SetController(controller)
	// Root auf Pages setzen (enthält LayoutView)
	if err := app.SetRoot(controller.GetLayout(), true).Run(); err != nil {
		panic(err)
	}
}
