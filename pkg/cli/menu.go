package cli

import (
	"log"

	"github.com/rivo/tview"
)

func Menu() {
	// Create a new application
	app := tview.NewApplication()

	// makes the menu, i'm planning to use a modular approach when it comes 
	list := tview.NewList().
		AddItem("Scan a Network", "", '1', nil).
		AddItem("About", "", '2', nil).
		AddItem("Exit", "Exit the program", '3', func() {
			app.Stop()
		})

	// Set the list as the root view and focus it
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		log.Fatalf("%+v", err)
	}
}
