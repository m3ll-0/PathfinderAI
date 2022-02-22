package main

type Board [][]string

type Node struct {
	level int
	parent *Node
	children []*Node
	currentPosition Coordinate
	goalPosition Coordinate
	board Board
	gValue float64
}

func (node *Node) generateChildren(){

	currentBoard := node.board

	coordinateMarkList := []Coordinate{}

	// Create list of coordinateMarks
	rightCoordinateMark := Coordinate{
		row:    node.currentPosition.row,
		column: node.currentPosition.column + 1,
	}

	leftCoordinateMark := Coordinate{
		row:    node.currentPosition.row,
		column: node.currentPosition.column - 1,
	}

	upperCoordinateMark := Coordinate{
		row:    node.currentPosition.row + 1,
		column: node.currentPosition.column,
	}

	lowerCoordinateMark := Coordinate{
		row:    node.currentPosition.row - 1,
		column: node.currentPosition.column,
	}

	coordinateMarkList = append(coordinateMarkList, rightCoordinateMark)
	coordinateMarkList = append(coordinateMarkList, leftCoordinateMark)
	coordinateMarkList = append(coordinateMarkList, upperCoordinateMark)
	coordinateMarkList = append(coordinateMarkList, lowerCoordinateMark)

	for _, coordinateMark := range coordinateMarkList{

		var newBoard = createCopy(currentBoard)

		// Mark current position as visited
		newBoard[node.currentPosition.row][node.currentPosition.column] = "o"

		// Check if coordinate is within board bounds
		if !coordinateMarkIsWithingBoardBounds(currentBoard, coordinateMark){
			continue
		}

		// Get new coordinate and check if it is free
		newMark := currentBoard[coordinateMark.row][coordinateMark.column]

		if newMark != "#" && newMark != "e" { // Coordinate is not a free space or finish
			continue
		}

		// Set player mark on new position
		newBoard[coordinateMark.row][coordinateMark.column] = "p"

		// Calculate and set g value
		totalTimesCoordinateHasBeenVisited := coordinateHeatMap[coordinateMark]
		gValue := totalTimesCoordinateHasBeenVisited

		childNode := &Node{
			level:           node.level + 1,
			parent:          node,
			currentPosition: coordinateMark,
			board:           newBoard,
			goalPosition:    node.goalPosition,
			gValue: float64(gValue),
		}

		node.children = append(node.children, childNode)
	}
}

func (nodeList PriorityQueue)removeNodeFromPriorityQueue(nodeToBeRemoved *Node) []*Node {

	newList := []*Node{}

	for _, node := range nodeList{
		if node != nodeToBeRemoved {
			newList = append(newList, node)
		}
	}

	return newList
}

func coordinateMarkIsWithingBoardBounds(board Board, coordinate Coordinate) bool{
	rowBound := len(board) - 1
	columnBound := len(board[0]) - 1

	if (!(coordinate.row <= rowBound)) || (!(coordinate.column <= columnBound) || (coordinate.row < 0) || (coordinate.column < 0)){
		return false
	}

	return true

}

func createCopy(currentBoard Board) Board{
	var newBoard = Board{}

	for _, row := range currentBoard{

		newRow := []string{}

		for _, cc := range row {
			newRow = append(newRow, cc)
		}

		copy(row, newRow)

		newBoard = append(newBoard, newRow)
	}

	return newBoard
}