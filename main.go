package main

import (
	"embed"
	"github.com/zserge/lorca"
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
var canvasMode = false
var currentBoard Board
const fillBoardThreads = 10

var updateBoardMethod UpdateBoardMethod
var generateBoardMethod GenerateBoardMethod
var fillBoardMethod FillBoardMethod

func main() {
	updateBoardMethod = updateTable
	generateBoardMethod = generateTableFromBoard
	fillBoardMethod = fillTableFromBoard

	loadBoardPresets()
	setupLorca()
}


