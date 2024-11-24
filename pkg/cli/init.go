package cli

import (
	"github.com/gdamore/tcell/v2"
)

// Initialize the tcell screen
func InitScreen() (tcell.Screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}
	return s, nil
}

// Clear the terminal screen
func ClearTerminal(s tcell.Screen) {
	s.Clear()
	s.Show()
}

