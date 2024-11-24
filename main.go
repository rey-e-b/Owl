package main

import (
	"github.com/rey-e-b/Owl/pkg/cli"
)

func main() {
	box := tview.NewBox().
	SetBorder(true).
	SetTitle("Owl")

	cli.Menu()
}
