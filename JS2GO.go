package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//*********************************************************************
//******** These methods are being called through JS functions ********
//*********************************************************************

func doPathFinding(solverMethodString string){

	stopSolver = false
	board := boardList[currentBoardCounter]
	currentPlayerPosition := getCoordinateFromBoardMark(board, "p")
	currentGoalPosition := getCoordinateFromBoardMark(board, "e")

	rootNode := &Node{board: board, currentPosition: currentPlayerPosition, goalPosition: currentGoalPosition}

	var solverMethod SolverMethod

	// Select solver method
	switch solverMethodString {
	case "aStar":
		solverMethod = aStar
		break
	case "BFS":
		solverMethod = BFS
		break
	case "DFS":
		solverMethod = DFS
		break
	default:
		println("Error: solved method unknown.")
		os.Exit(1)
	}

	resetStatistics() // Reset statistics for visual purposes
	timeStart = time.Now()
	nodesVisitedCounter = 0

	// Reset the board each time pathfinding is started using JS to speed up reload
	loadCurrentBoard()

	// Start solver
	solvedNode := solverMethod(rootNode)

	// Load current board again to prepare for backtracking
	loadCurrentBoard()

	// Start backtracking from solution to root node
	backtrack(solvedNode)

	setStatistics(nodesVisitedCounter, fmt.Sprint(time.Since(timeStart)), solvedNode.level)

	coordinatesVisited = []Coordinate{} // todo: ugly hack to pass variable function, needed for DFS
}

func stopSolvers(){
	stopSolver = true
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
	saveCurrentBoard() // After filling the board, save the current board in JS for reloading purposes


	resetStatistics()
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