package cli

import "fmt"

// PrintMenu displays the options in the menu with the selected option highlighted
func PrintMenu(options []string, selectedIndex int) {
	fmt.Println("Owl Menu")
	fmt.Println()
	for i, option := range options {
		if i == selectedIndex {
			// Highlight selected option
			fmt.Printf("[*] %s\n", option)
		} else {
			fmt.Printf("[ ] %s\n", option)
		}
	}
}

