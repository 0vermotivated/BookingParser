package main

import (
	"BookingParser/internal/GUI"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := GUI.SetupMainWindow(app)
	window.ShowAndRun()
}
