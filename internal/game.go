package ui

import (
	board "github.com/MiguelCiulog/PicGo/internal/ui"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type model struct {
	board     board.Model
	stopwatch stopwatch.Model

	gameEnded bool

	err        error
	mouseEvent tea.MouseEvent
}

func (m model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return m, tea.Quit
		}

	case tea.MouseMsg:
		m.mouseEvent = tea.MouseEvent(msg)
	}

	return m, nil
}

func (m model) View() string {
	return m.stopwatch.View()
	// style := lipgloss.NewStyle().
	// 	BorderStyle(lipgloss.NormalBorder()).
	// 	Background(lipgloss.Color("#121212"))

	// stuff := "uwu, owo"
	// mainView := style.Render(stuff)
	// return mainView
}

func NewModel() model {
	return model{
		board:      board.NewModel(5, 5),
		stopwatch:  stopwatch.Model{},
		gameEnded:  false,
		err:        nil,
		mouseEvent: tea.MouseEvent{},
	}
}

//     mdl := model{
//     	board:      board.Model{},
//     	stopwatch:  stopwatch.Model{},
//     	gameEnded:  false,
//     	err:        nil,
//     	mouseEvent: tea.MouseEvent{},
//     }
// 	game := GenerateGrid(5, 5)
// 	return game
// }
