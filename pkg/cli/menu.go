package cli

import (
	"github.com/gdamore/tcell/v2"
)

// Print the menu with selection highlighting
func PrintMenu(s tcell.Screen, options []string, selectedIndex int) {
	width, height := s.Size()

	for i, option := range options {
		style := tcell.StyleDefault
		if i == selectedIndex {
			style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite) // highlight selected option
		} else {
			style = style.Foreground(tcell.ColorWhite) // default style for other options
		}

		// print each option centered
		for j, r := range option {
			s.SetContent(width/2-len(option)/2+j, height/2-len(options)/2+i, r, nil, style)
		}
	}

	s.Show()
}

