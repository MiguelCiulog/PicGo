package main

import (
	"fmt"
)

func main() {
	fmt.Println("uwu")
}

const (
	blankGuess     = 0
	correctGuess   = 1
	incorrectGuess = 2
	filledSquare   = "■"
	blankSquare    = "□"
)

type Grid struct {
	Rows   []int
	Column []int

	Clues map[*int][]int
}


func createBlankGrid(x int, y int) Grid {
	gr := make([][]int, x, y)
	return gr
}

func play() {
	fmt.Println("Start gaming")
}
