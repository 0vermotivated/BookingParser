package main

import (
	"BookingParser/internal/GUI"
)

func main() {
	app := app.New()
	window := GUI.SetupMainWindow(app)
	window.ShowAndRun()
}
