package cli

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"fmt"
)

// ReadInput handles user key presses
func ReadInput() string {
	app := tview.NewApplication()
	inputCh := make(chan string)

	// set up the key listener
	go func() {
		for {
			// set up input listener
			switch ev := app.GetEvent(); ev.(type) {
			case *tcell.EventKey:
				key := ev.(*tcell.EventKey).Key()
				switch key {
				case tcell.KeyArrowUp: // for up arrow
					inputCh <- "w"
				case tcell.KeyArrowDown: // for down arrow
					inputCh <- "s"
				case tcell.KeyEnter: // for enter
					inputCh <- "enter"
				}
			}
		}
	}()

	return <-inputCh
}

