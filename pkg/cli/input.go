package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ClearTerminal clears the terminal screen
func ClearTerminal() {
	cmd := exec.Command("clear") // use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ReadInput reads user input from the terminal
func ReadInput() string {
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input) // convert input to lowercase for consistency
	if input == "" {
		input = "enter" // default to enter if no input
	}
	return input
}

