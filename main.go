package main

import (
	"fmt"
	"log"
	"time"
	"github.com/rey-e-b/Owl/pkg/cli"
)

func main() {
	options := []string{"Scan Network", "About", "Exit"}
	selectedIndex := 0

	// initialize screen with tcell
	s, err := cli.InitScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	for {
		// clear terminal and print menu
		cli.ClearTerminal(s)
		cli.PrintMenu(s, options, selectedIndex)

		// handle user input
		input := cli.ReadInput(s)

		// handle menu navigation and selection
		switch input {
		case "up": // move up
			if selectedIndex > 0 {
				selectedIndex--
			}
		case "down": // move down
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

