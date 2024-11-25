package cli

import (
	"log"
	"github.com/rivo/tview"
)

// menu defines the parent list
func Menu(app *tview.Application) error {
	list := tview.NewList().
		AddItem("Scan a Network", "", '1', func() {
			// switch to the second menu
			SecondMenu(app)
		}).
		AddItem("About", "", '2', nil).
		AddItem("Exit", "Exit the program", '3', func() {
			app.Stop()
		})

	// set the root view and focus the list
	return app.SetRoot(list, true).SetFocus(list).Run()
}

