package main

import (
	"embed"
	"fmt"
	"github.com/zserge/lorca"
	"os"
	"time"
)

//go:embed www
var fs embed.FS
var ui, err = lorca.New("", "", 1000, 1000)
var currentBoardCounter = -1
var timeStart = time.Now()
var nodesVisitedCounter = 0
var stopSolver = false
var boardList = []Board{}
var speedMode = true

func main() {
	setupLorca()
}

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
	case "aStarCancerous":
		solverMethod = aStarCancerous
		break
	case "DFS":
		solverMethod = DFS
		break
	default:
		println("Error: solved method unknown.")
		os.Exit(1)
		break
	}

	// Reset statistics
	resetStatistics()
	timeStart = time.Now()
	nodesVisitedCounter = 0

	solvedNode := solverMethod(rootNode)

	setStatistics(nodesVisitedCounter, fmt.Sprint(time.Since(timeStart)), solvedNode.level)

	coordinatesVisited = []Coordinate{} // todo: ugly hack to pass variable function, needed for DFS
}



