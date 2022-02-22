package main

import (
	"fmt"
)

func resetStatistics(){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v');", "?", "?"))
}

func setStatistics(amountOfNodesVisited int, totalAmountOfTime string){
	ui.Eval(fmt.Sprintf("setStatistics('%v', '%v');", amountOfNodesVisited, totalAmountOfTime))
}

func generateGridFromBoard(board Board){

	// Get dimensions
	rowDimension := len(board)
	columnDimension := len(board[0])

	// First generate grid and assign cell values coordinates
	ui.Eval(fmt.Sprintf("generateGrid(%v,%v);", rowDimension, columnDimension))
}

func fillGridFromBoard(board Board){
	for rowCounter, row := range board{
		for columnCounter, column := range row {

			if column == "#"{
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "#5a5757"))
			} else if column == "*" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "red"))
			} else if column == "p" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "green"))
			}  else if column == "e" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "black"))
			} else if column == "o" {
				ui.Eval(fmt.Sprintf("setCellBGColor(%v,%v,'%v');", rowCounter, columnCounter, "lightgreen"))
			}

		}
	}
}
