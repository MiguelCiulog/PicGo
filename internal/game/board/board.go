package board

import (
	// "fmt"
	"math/rand"
)

type cellStatus int

const (
	blank cellStatus = iota
	filled
	crossed
)

type clue struct {
	hint   []int
	status []bool
}

type Model struct {
	board       [][]cellStatus
	columnClues []clue
	rowClues    []clue

	maxRowCells, maxColumnCells int
}

func (model *Model) isSegmentSolved(segment []clue) bool {
	for _, cell := range segment {
		for _, isFullfilled := range cell.status {
			if !isFullfilled {
				return false
			}
		}
	}
	return true
}

func (model *Model) IsBoardSolved() bool {
	areColumnsSolved := model.isSegmentSolved(model.columnClues)
	if !areColumnsSolved {
		return false
	}

	areRowsSolved := model.isSegmentSolved(model.rowClues)
	if !areRowsSolved {
		return false
	}

	return true
}

func NewModel(maxRowCells, maxColumnCells int) (model Model) {
	model.maxRowCells = maxRowCells
	model.maxColumnCells = maxColumnCells

	model.board = model.createBlankBoard()
	model.rowClues, model.columnClues = model.getNonogramHints()

	// fmt.Println(model)
	return model
}

// Create a 2D array of booleans with default blank cell status
func (model *Model) createBlankBoard() [][]cellStatus {
	var board [][]cellStatus
	for i := 0; i < model.maxRowCells; i++ {
		row := make([]cellStatus, model.maxColumnCells)
		board = append(board, row)
	}
	return board
}

// Create a random board with filled stuff
func (model *Model) createRandomBoard() [][]cellStatus {
	var board [][]cellStatus

	// TODO: Create an algorithm for symmetry
	for i := 0; i < model.maxColumnCells; i++ { // Columns
		cellHolder := make([]cellStatus, 0)
		for j := 0; j < model.maxRowCells; j++ { // Rows
			cell := blank
			if shouldBeFilled() {
				cell = filled
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

func generateSegmentHints(
	randomBoard [][]cellStatus,
	mainSegmentSize int,
	secondarySegmentSize int,
) []clue {
	clue := make([]clue, mainSegmentSize)

	// Process segment (rows then columns or viceversa)
	for i := 0; i < mainSegmentSize; i++ {
		row := randomBoard[i]
		hint := []int{}
		status := []bool{}
		count := 0

		for j := 0; j < secondarySegmentSize; j++ {
			if row[j] == filled {
				count++
			} else {
				if count > 0 {
					hint = append(hint, count)
					status = append(status, false)
					count = 0
				}
			}
		}

		if count > 0 {
			hint = append(hint, count)
			status = append(status, false)
		}

		clue[i].hint = hint
		clue[i].status = status
	}
	return clue
}

func (model *Model) getNonogramHints() (rowClues, colClues []clue) {
	randomBoard := model.createRandomBoard()
	numRows := len(randomBoard)
	numCols := len(randomBoard[0])

	rowClues = generateSegmentHints(randomBoard, numRows, numCols)
	colClues = generateSegmentHints(randomBoard, numCols, numRows)

	return rowClues, colClues
}
