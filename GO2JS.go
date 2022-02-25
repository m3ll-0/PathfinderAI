package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//*************************************************
//******** These methods call JS functions ********
//*************************************************

func alertMessage(message string){
	ui.Eval(fmt.Sprintf("alert('%v');", message))
}

func resetStatistics(){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v', '%v');", "?", "?", "?"))
}

func setStatistics(totalAmountOfNodesVisited int, totalAmountOfTime string, totalPathLength int){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v', '%v');", totalAmountOfNodesVisited, totalAmountOfTime, totalPathLength))
}

//******** Table functions ********

func generateCanvasFromBoard(board Board){

	// Get dimensions
	rowDimension := len(board)
	columnDimension := len(board[0])

	// First generate grid and assign cell values coordinates
	ui.Eval(fmt.Sprintf("generateCanvas(%v,%v);", rowDimension, columnDimension))
}

func updateTable(prevPlayerCoordinate Coordinate, curPlayerCoordinate Coordinate){
	ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", prevPlayerCoordinate.row, prevPlayerCoordinate.column, "#2383d6"))
	ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", curPlayerCoordinate.row, curPlayerCoordinate.column, "#07418a"))
}

type cell struct {
	coord Coordinate
	mark string
}

var wg sync.WaitGroup

func fillTableFromBoard(board Board){

	var flatCellList []cell

	// Create flat list of cells
	for rowCounter, row := range board{
		for columnCounter, column := range row {
			flatCellList = append(flatCellList, cell{coord: Coordinate{row: rowCounter, column: columnCounter}, mark: column})
		}
	}

	for i := range flatCellList {
		j := rand.Intn(i + 1)
		flatCellList[i], flatCellList[j] = flatCellList[j], flatCellList[i]
	}

	// Create a list with length 10 of list of cells
	cellListList := [fillBoardThreads][]cell{}

	// Divide cells over the 10 lists so that they can be run as separate threads
	for cellNumber, cell := range flatCellList{
		cellListList[cellNumber%fillBoardThreads] = append(cellListList[cellNumber%fillBoardThreads], cell)
	}

	sliceLength := fillBoardThreads // numberOfThreads
	wg.Add(sliceLength)

	for _, cellList := range cellListList {
		go threadFillCell(cellList) // Start thread for each list of cells
	}

	wg.Wait()
}

func threadFillCell(cells []cell){
	for _, cell := range cells{
		//println(fmt.Sprint(cell))
		fillCell(cell.mark, cell.coord.row, cell.coord.column)
	}
	defer wg.Done()
}

func fillCell(mark string, row int, column int){
	if mark == "#"{
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", row, column, "#5a5757"))
	} else if mark == "*" {
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", row, column, "red"))
	} else if mark == "p" {
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", row, column, "#07418a"))
	}  else if mark == "e" {
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", row, column, "green"))
	} else if mark == "o" {
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", row, column, "#2383d6"))
	}
}

//******** Canvas functions ********

func generateTableFromBoard(board Board){

	// Get dimensions
	rowDimension := len(board)
	columnDimension := len(board[0])

	// First generate grid and assign cell values coordinates
	ui.Eval(fmt.Sprintf("generateGrid(%v,%v);", rowDimension, columnDimension))
}

func updateCanvas(prevPlayerCoordinate Coordinate, curPlayerCoordinate Coordinate){
	ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", prevPlayerCoordinate.row, prevPlayerCoordinate.column, "#2383d6"))
	ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", curPlayerCoordinate.row, curPlayerCoordinate.column, "#07418a"))
}

func fillCanvasFromBoard(board Board){
	for rowCounter, row := range board{
		for columnCounter, column := range row {

			if column == "#"{
				ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", rowCounter, columnCounter, "#5a5757"))
			} else if column == "*" {
				ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", rowCounter, columnCounter, "red"))
			} else if column == "p" {
				ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", rowCounter, columnCounter, "#07418a"))
			}  else if column == "e" {
				ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", rowCounter, columnCounter, "green"))
			} else if column == "o" {
				ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", rowCounter, columnCounter, "#2383d6"))
			}

		}
	}
}

func backtrack(node *Node){

	// Set goal
	ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", node.currentPosition.row, node.currentPosition.column, "#07418a"))
	for true {

		node = node.parent

		if node.parent == nil{
			// Set starting node
			ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", node.currentPosition.row, node.currentPosition.column, "#2383d6"))
			break
		}

		// Set intermediate node
		ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", node.currentPosition.row, node.currentPosition.column, "#2383d6"))
	}

}

func loadCurrentBoard(){
	ui.Eval("loadCurrentBoard()")
}

func saveCurrentBoard(){
	ui.Eval("saveCurrentBoard()")
}