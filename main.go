package main

import (
	"fmt"
	"harpoon/panel"
	"os"
)

func main() {
	if err := panel.NewGUI().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
