package main

// A simple program that opens the alternate screen buffer and displays mouse
// coordinates and events.

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
    game "github.com/MiguelCiulog/PicGo/internal/game"
)

func main() {
	p := tea.NewProgram(game.NewModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

