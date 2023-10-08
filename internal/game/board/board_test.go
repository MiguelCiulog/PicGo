package board

import (
	"testing"
)

func TestBoardCreation(t *testing.T) {
	model := NewModel(5, 5)

	t.Logf("%#v\n", model)
}

func TestIfBoardIsSolved(t *testing.T) {
	unsolvedBoard := Model{
		board: [][]cellStatus{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		columnClues: []clue{
			{hint: []int{1, 1}, status: []bool{false, false}},
			{hint: []int{2}, status: []bool{false}},
			{hint: []int{1, 1}, status: []bool{false, false}},
			{hint: []int{}, status: []bool{}},
			{hint: []int{3}, status: []bool{false}},
		},
		rowClues: []clue{
			{hint: []int{1, 1}, status: []bool{false, false}},
			{hint: []int{2}, status: []bool{false}},
			{hint: []int{1, 1}, status: []bool{false, false}},
			{hint: []int{}, status: []bool{}},
			{hint: []int{3}, status: []bool{false}},
		},
		maxRowCells:    5,
		maxColumnCells: 5,
	}
	if unsolvedBoard.IsBoardSolved() {
		t.Log(unsolvedBoard.board)
		t.Error("Unsolved board shows as solved.")
	}

	solvedBoard := unsolvedBoard
	solvedBoard.rowClues = []clue{
		{hint: []int{1, 1}, status: []bool{true, true}},
		{hint: []int{2}, status: []bool{true}},
		{hint: []int{1, 1}, status: []bool{true, true}},
		{hint: []int{}, status: []bool{}},
		{hint: []int{3}, status: []bool{true}},
	}
	solvedBoard.columnClues = []clue{
		{hint: []int{1, 1}, status: []bool{true, true}},
		{hint: []int{2}, status: []bool{true}},
		{hint: []int{1, 1}, status: []bool{true, true}},
		{hint: []int{}, status: []bool{}},
		{hint: []int{3}, status: []bool{true}},
	}

	if !solvedBoard.IsBoardSolved() {
		t.Log(solvedBoard.board)
		t.Error("Solved board shows as unsolved.")
	}
}
