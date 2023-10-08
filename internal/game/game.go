package ui

import (
	board "github.com/MiguelCiulog/PicGo/internal/game/board"
	"github.com/MiguelCiulog/PicGo/internal/game/styles"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	keys "github.com/MiguelCiulog/PicGo/internal/game/keys"
	zone "github.com/lrstanley/bubblezone"
)

type model struct {
	board     board.Model
	stopwatch stopwatch.Model

	selectedCell [][]int

	gameEnded     bool
	width, height int

	err        error
	mouseEvent tea.MouseEvent
	keyEvent   keys.KeyMap
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

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// case tea.MouseMsg:
		// 	m.mouseEvent = tea.MouseEvent(msg)
	}

	return m, nil
}

func (m model) View() string {
	// fmt.Println("%#v\n", m.board)
	// return zone.Scan(styles.BlankCell.Render("x","x"))
	// var out strings.Builder

	uwu := styles.BlankCell.Render(" ")
	owo := styles.SelectedCell.Render(" ")
	// x := lipgloss.JoinVertical()
	x := lipgloss.JoinHorizontal(lipgloss.Center, uwu, owo, uwu)
	// out.WriteString(uwu)
	// out.WriteString(owo)
	// out.WriteString(uwu)

	return zone.Scan(lipgloss.PlaceHorizontal(m.width, lipgloss.Center, x))
	// return zone.Scan(lipgloss.JoinHorizontal(lipgloss.Center, uwu,owo,uwu))
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
