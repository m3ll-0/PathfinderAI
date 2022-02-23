package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printBoard(board Board){
	println("")
	for _, row := range board {
		println(fmt.Sprint(row))
	}
}

func printBoardCopy(board Board){
	println("\n")
	print("Board{\n")
	for _, row := range board {
		print("{")
		for _, column := range row {
			print("\""+ column +"\", ")
		}
		print("},\n")
	}
	print("}")
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
	if currentBoardCounter >= len(boardList) - 1 {
		alertMessage("Error: No boards left.")
		return
	}

	currentBoardCounter++

	board := boardList[currentBoardCounter]
	generateBoardMethod(board)

	currentBoard = board // Set current board, so you can switch between canvas mode and table mode intermittently
	fillBoardMethod(board)
	resetStatistics()
}

func stopSolvers(){
	stopSolver = true
}

func generateRandomLevel(){

	rand.Seed(time.Now().UnixNano())
	randAmountOfRows := rand.Intn(80 - 5 + 1) + 5
	randAmountOfColumns := rand.Intn(210 - 5 + 1) + 5

	var board Board

	for i := 0; i <= randAmountOfRows; i++ {

		var boardRow []string

		for j := 0; j <= randAmountOfColumns; j++ {

			randMarkNumber := rand.Intn(10)

			mark := ""

			// Determine ratio
			if randMarkNumber <= 6 {
				mark = "#"
			} else {
				mark = "*"
			}

			boardRow = append(boardRow, mark)
		}
		board = append(board, boardRow)
	}

	// Set random player & goal
	randomPlayerRow := rand.Intn(randAmountOfRows)
	randomPlayerColumn := rand.Intn(randAmountOfColumns)
	randomGoalRow := rand.Intn(randAmountOfRows)
	randomGoalColumn := rand.Intn(randAmountOfColumns)

	board[randomPlayerRow][randomPlayerColumn] = "p"
	board[randomGoalRow][randomGoalColumn] = "e"

	printBoardCopy(board)

	boardList = append(boardList[:currentBoardCounter+1], boardList[currentBoardCounter:]...)
	boardList[currentBoardCounter] = board
	currentBoardCounter--

	loadNextBoard()
}

func toggleSpeedMode(){
	speedMode = !speedMode
}

func toggleCanvasMode(){

	canvasMode = !canvasMode

	// Reset listeners
	if canvasMode {
		updateBoardMethod = updateCanvas
		generateBoardMethod = generateCanvasFromBoard
		fillBoardMethod = fillCanvasFromBoard
	} else {
		updateBoardMethod = updateTable
		generateBoardMethod = generateTableFromBoard
		fillBoardMethod = fillTableFromBoard
	}

	// Reload current board
	generateBoardMethod(currentBoard)
	fillBoardMethod(currentBoard)



}