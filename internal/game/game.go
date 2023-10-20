package ui

import (
	"strconv"

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
	boardStr := make([]string, m.board.MaxRowCells)
	for x_idx, row := range m.board.Board {
		rowStr := make([]string, m.board.MaxColumnCells)

		for y_idx, cell := range row {
			var style lipgloss.Style
			switch cell {
			case board.Blank:
				if (x_idx+y_idx)%2 == 0 {
					style = styles.BlankCellVar1
				} else {
					style = styles.BlankCellVar2

				}
			case board.Filled:
				style = styles.FilledCell
			case board.Crossed:
				style = styles.CrossedCell
			}
			coordinates := strconv.Itoa(x_idx) + strconv.Itoa(y_idx)
			tmpStr := zone.Mark(coordinates, style.Render(" "))
			rowStr = append(rowStr, tmpStr)
		}

		boardStr = append(boardStr, lipgloss.JoinHorizontal(lipgloss.Center, rowStr...))
	}

	finalstr := lipgloss.JoinVertical(lipgloss.Center, boardStr...)
	return zone.Scan(lipgloss.PlaceHorizontal(m.width, lipgloss.Center, finalstr))
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
