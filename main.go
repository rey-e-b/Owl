package main

import (
	"log"

	"github.com/rey-e-b/Owl/pkg/cli"
	"github.com/rivo/tview"
)

func main() {
	// create a new application
	app := tview.NewApplication()

	// call the main menu and pass the app instance
	if err := cli.Menu(app); err != nil {
		log.Fatalf("error running app: %+v", err)
	}
}

