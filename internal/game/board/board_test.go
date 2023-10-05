package board

import (
	"testing"
)

func TestIfBoardIsSolved(t *testing.T) {
	unsolvedBoard := Model{
		// 5x5 board
		board: [][]cell{
			{
				{
					status:         blank,
					shouldBeFilled: false,
					modifiable:     true,
				}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true},
				{
					status:         blank,
					shouldBeFilled: false,
					modifiable:     true,
				}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true},
				{
					status:         blank,
					shouldBeFilled: true,
					modifiable:     true,
				}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true},
				{
					status:         blank,
					shouldBeFilled: true,
					modifiable:     true,
				}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true},
				{
					status:         blank,
					shouldBeFilled: false,
					modifiable:     true,
				}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true},
			},
		},
		columnClues: [][]int{{4}, {2}, {2, 2}, {5}, {1}},
		rowClues:    [][]int{{4}, {2}, {2, 2}, {5}, {1}},
	}
	if unsolvedBoard.IsBoardSolved() {
		t.Log(unsolvedBoard.board)
		t.Error("Unsolved board shows as solved.")
	}

	solvedBoard := Model{
		// 5x5 board
		board: [][]cell{
			{
				{
					status:         blank,
					shouldBeFilled: false,
					modifiable:     true,
				}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}, {status: filled, shouldBeFilled: true, modifiable: true}, {status: blank, shouldBeFilled: false, modifiable: true}},
		},
		columnClues: [][]int{{4}, {2}, {2, 2}, {5}, {1}},
		rowClues:    [][]int{{4}, {2}, {2, 2}, {5}, {1}},
	}
	if !solvedBoard.IsBoardSolved() {
		t.Log(solvedBoard.board)
		t.Error("Solved board shows as unsolved.")
	}
}
