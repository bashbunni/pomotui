package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charlieroth/pomotui/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	model := model.New()
	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
