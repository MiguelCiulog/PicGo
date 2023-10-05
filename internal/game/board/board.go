package board

import (
	"fmt"
	"math/rand"
)

type cellStatus int

const (
	blank cellStatus = iota
	filled
	crossed
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

func (model *Model) IsBoardSolved() bool {
	board := model.board

	for _, row := range board {
		for _, cell := range row {
			switch cell.status {
			case filled:
				if cell.shouldBeFilled == false {
					return false
				}
			case blank:
				if cell.shouldBeFilled == true {
					return false
				}
			case crossed:
				if cell.shouldBeFilled == true {
					return false
				}
			}
		}
	}
	return true
}

func NewModel(maxWidth int, maxHeight int) (model Model) {
	board := model.createRandomBoard(maxWidth, maxHeight)
	model.board = board

	rowHints, colHints := model.getNonogramHints()
	model.rowClues = rowHints
	model.columnClues = colHints

	fmt.Println(model)
	return model
}

func (model *Model) createRandomBoard(maxWidth int, maxHeight int) [][]cell {
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

func (model *Model) generateSegmentHints(mainSegmentSize, secondarySegmentSize int) [][]int {
	board := model.board
	hints := make([][]int, mainSegmentSize)

	// Process segment (rows then columns or viceversa)
	for i := 0; i < mainSegmentSize; i++ {
		row := board[i]
		hint := []int{}
		count := 0

		for j := 0; j < secondarySegmentSize; j++ {
			if row[j].shouldBeFilled {
				count++
			} else {
				if count > 0 {
					hint = append(hint, count)
					count = 0
				}
			}
		}

		if count > 0 {
			hint = append(hint, count)
		}

		hints[i] = hint
	}
	return hints
}

func (model *Model) getNonogramHints() (rowHints, colHints [][]int) {
	board := model.board
	numRows := len(board)
	numCols := len(board[0])

	rowHints = model.generateSegmentHints(numRows, numCols)
	colHints = model.generateSegmentHints(numCols, numRows)

	return rowHints, colHints
}

// func (m Model) Init() tea.Cmd {
// 	return nil
// }

// func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
// 			return m, tea.Quit
// 		}

// 	case tea.MouseMsg:
// 		//
// 	}

// 	return m, nil
// }

// func (m Model) View() string {
// 	style := lipgloss.NewStyle().
// 		BorderStyle(lipgloss.NormalBorder()).
// 		Background(lipgloss.Color("#121212"))

// 	stuff := "uwu, owo"
// 	mainView := style.Render(stuff)
// 	return mainView
// }
