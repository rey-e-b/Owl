package cli

import (
	"github.com/rivo/tview"
)

// Menu creates a menu inside a frame with a border
func Menu(app *tview.Application) error {
	// Create the list (menu items)
	list := tview.NewList().
		AddItem("Scan a Network", "", '1', func() {
			SecondMenu(app)	
		}).
		AddItem("About", "", '2', nil).
		AddItem("Exit", "Exit the program", '3', func() {
			app.Stop()
		})


	// Create a box with a border around the list
	list.SetBorder(true).SetTitle(" Main Menu ")

	// Set the list as the root view and run the app
	return app.SetRoot(list, true).SetFocus(list).Run()
}

