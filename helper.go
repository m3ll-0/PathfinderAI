package main

import "fmt"

func printBoard(board Board){
	println("")
	for _, row := range board {
		println(fmt.Sprint(row))
	}
}

func getCoordinateFromBoardMark(board Board, mark string) Coordinate{
	for rowCounter, row := range board{
		for colCounter, col := range row {
			if col == mark{
				return Coordinate{row: rowCounter, column: colCounter}
			}
		}
	}

	return Coordinate{row: 0, column: 0}
}

func loadNextBoard() {
	currentBoardCounter++

	boards := getBoards()
	board := boards[currentBoardCounter]
	generateGridFromBoard(board)
	fillGridFromBoard(board)
	resetStatistics()
}