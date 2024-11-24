package cli

import (
	"github.com/gdamore/tcell/v2"
)

// Read input from the user
func ReadInput(s tcell.Screen) string {
	event := s.PollEvent()
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyArrowUp:
			return "up"
		case tcell.KeyArrowDown:
			return "down"
		case tcell.KeyEnter:
			return "enter"
		}
	}
	return ""
}

