package main

import (
	"fmt"
	"os"
	"time"
	"Owl/pkg/cli"
)

func main() {
	options := []string{"Scan Network", "About", "Exit"}
	selectedIndex := 0

	for {
		// clear terminal and print menu
		cli.ClearTerminal()
		cli.PrintMenu(options, selectedIndex)

		// handle user input
		input := cli.ReadInput()

		// handle menu navigation and selection
		switch input {
		case "w": // move up
			if selectedIndex > 0 {
				selectedIndex--
			}
		case "s": // move down
			if selectedIndex < len(options)-1 {
				selectedIndex++
			}
		case "enter": // select option
			handleSelection(selectedIndex)
			if selectedIndex == 2 { // exit option
				return
			}
		}
	}
}

// handle selection and perform action based on selected option
func handleSelection(selectedIndex int) {
	switch selectedIndex {
	case 0:
		fmt.Println("Scanning Network...") // placeholder
	case 1:
		fmt.Println("About Owl...") // placeholder
	case 2:
		fmt.Println("Exiting...")
	}
	time.Sleep(1 * time.Second) // simulate delay before returning to menu
}

