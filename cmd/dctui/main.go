package main

import (
	service "dctui/internal/services"
	"dctui/internal/ui"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	header := ui.NewHeaderView()
	commandBar := ui.NewCommandBarView(app)
	layout := ui.NewLayoutView(header, commandBar, app)
	docker := service.NewDockerService()
	controller := ui.NewUIController(app, layout, docker)
	//controller.AddPage("projects", ui.NewProjectListPage())
	controller.AddPage("projects", ui.NewProjectOverview(app))
	controller.AddPage("projectView", ui.NewProjectView(app))
	controller.SwitchToPage("projects", docker.FetchProjects())

	commandBar.SetController(controller)

	if err := app.SetRoot(layout.GetPrimitive(), true).Run(); err != nil {
		panic(err)
	}
}
