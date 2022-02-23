package main

import (
	"fmt"
)

func alertMessage(message string){
	ui.Eval(fmt.Sprintf("alert('%v');", message))
}

func resetStatistics(){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v', '%v');", "?", "?", "?"))
}

func setStatistics(totalAmountOfNodesVisited int, totalAmountOfTime string, totalPathLength int){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v', '%v');", totalAmountOfNodesVisited, totalAmountOfTime, totalPathLength))
}

func generateTableFromBoard(board Board){

	// Get dimensions
	rowDimension := len(board)
	columnDimension := len(board[0])

	// First generate grid and assign cell values coordinates
	ui.Eval(fmt.Sprintf("generateGrid(%v,%v);", rowDimension, columnDimension))
}

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

func updateCanvas(prevPlayerCoordinate Coordinate, curPlayerCoordinate Coordinate){
	ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", prevPlayerCoordinate.row, prevPlayerCoordinate.column, "#2383d6"))
	ui.Eval(fmt.Sprintf("updateCanvasCell(%v,%v,'%v');", curPlayerCoordinate.row, curPlayerCoordinate.column, "#07418a"))
}

func fillTableFromBoard(board Board){
	for rowCounter, row := range board{
		for columnCounter, column := range row {
			if column == "*" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "red"))
			} else if column == "p" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", rowCounter, columnCounter, "#07418a"))
			}  else if column == "e" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", rowCounter, columnCounter, "green"))
			} else if column == "o" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v', true);", rowCounter, columnCounter, "#2383d6"))
			}

		}
	}
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
