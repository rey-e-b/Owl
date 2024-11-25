package cli

import (
	"github.com/rivo/tview"
)

// secondmenu displays the second list
func SecondMenu(app *tview.Application) {
	list := tview.NewList().
		AddItem("Suboption 1", "", '1', nil).
		AddItem("Suboption 2", "", '2', nil).
		AddItem("Back to Main Menu", "", 'b', func() {
			// switch back to the main menu
			Menu(app)
		})
	list.SetBorder(true).SetTitle(" Sub Menu ")
	
	// set the second menu as the root
	app.SetRoot(list, true).SetFocus(list)
}

