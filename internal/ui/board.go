package board

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type cellStatus int

const (
	blank cellStatus = iota
	incorrect
	correct
)

type cell struct {
	status         cellStatus
	shouldBeFilled bool
	modifiable     bool
}

type Model struct {
	board       [][]cell
	columnClues [][]int
	rowClues    [][]int
}

func NewModel(maxWidth int, maxHeight int) Model {
	board := createRandomBoard(maxWidth, maxHeight)
	rowHints, colHints := getNonogramHints(board)

	model := Model{
		board:       board,
		columnClues: colHints,
		rowClues:    rowHints,
	}
    fmt.Println(model)
	return model
}

func createRandomBoard(maxWidth int, maxHeight int) [][]cell {
	var board [][]cell

	for i := 0; i < maxWidth; i++ { // Columns
		cellHolder := make([]cell, 0)
		for j := 0; j < maxHeight; j++ { // Rows
			cell := cell{
				status:     blank,
				modifiable: true,
			}
			if shouldBeFilled() {
				cell.shouldBeFilled = true
			} else {
				cell.shouldBeFilled = false
			}

			cellHolder = append(cellHolder, cell)
		}
		board = append(board, cellHolder)
	}
	return board
}

func shouldBeFilled() bool {
	return rand.Intn(2) == 1
}

func getHint(board [][]cell, mainSegmentSize int, secondarySegmentSize int) [][]int {
	rowHints := make([][]int, mainSegmentSize)
	// Process rows
	for i := 0; i < mainSegmentSize; i++ {
		row := board[i]
		rowHint := []int{}
		count := 0

		for j := 0; j < secondarySegmentSize; j++ {
			if row[j].shouldBeFilled {
				count++
			} else {
				if count > 0 {
					rowHint = append(rowHint, count)
					count = 0
				}
			}
		}

		if count > 0 {
			rowHint = append(rowHint, count)
		}

		rowHints[i] = rowHint
	}
	return rowHints
}

func getNonogramHints(board [][]cell) (rowHints, colHints [][]int) {
	numRows := len(board)
	numCols := len(board[0])

	// Initialize the slices to store row and column hints
	rowHints = make([][]int, numRows)
	colHints = make([][]int, numCols)

	// Process rows
	for i := 0; i < numRows; i++ {
		row := board[i]
		rowHint := []int{}
		count := 0

		for j := 0; j < numCols; j++ {
			if row[j].shouldBeFilled {
				count++
			} else {
				if count > 0 {
					rowHint = append(rowHint, count)
					count = 0
				}
			}
		}

		if count > 0 {
			rowHint = append(rowHint, count)
		}

		rowHints[i] = rowHint
	}

	// Process columns
	for j := 0; j < numCols; j++ {
		col := make([]bool, numRows)
		colHint := []int{}
		count := 0

		for i := 0; i < numRows; i++ {
			col[i] = board[i][j].shouldBeFilled
		}

		for i := 0; i < numRows; i++ {
			if col[i] {
				count++
			} else {
				if count > 0 {
					colHint = append(colHint, count)
					count = 0
				}
			}
		}

		if count > 0 {
			colHint = append(colHint, count)
		}

		colHints[j] = colHint
	}

	return rowHints, colHints
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return m, tea.Quit
		}

	case tea.MouseMsg:
		//
	}

	return m, nil
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Background(lipgloss.Color("#121212"))

	stuff := "uwu, owo"
	mainView := style.Render(stuff)
	return mainView
}

