package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

var coordinatesVisited []Coordinate

func DFS(node *Node) *Node {

	// Check if user cancelled
	if stopSolver {
		return node
	}

	nodesVisitedCounter++

	// Update current board in browser
	if node.parent != nil {
		if speedMode {
			updateBoard(node.currentPosition, node.parent.currentPosition)
		} else {
			fillGridFromBoard(node.board)
		}
	}

	// Return node if goal is reached
	if node.currentPosition == node.goalPosition {
		fillGridFromBoard(node.board)
		return node
	}

	var result *Node

	skip := false

	// Check if coordinate has already been visited
	for _, visitedCoordinate := range coordinatesVisited {
		if visitedCoordinate == node.currentPosition {
			skip = true
		}
	}

	if !skip {

		coordinatesVisited = append(coordinatesVisited, node.currentPosition)

		node.generateChildren()

		// For every child, call dfs
		for _, cn := range node.children{
			result = DFS(cn)

			// When node is being return from head, result is set. Result first has to be checked whether it is nul, then check goal
			//and return node so that node can be returned from recursive function
			if result != nil {
				// Return node if goal is reached
				if result.currentPosition == result.goalPosition {
					break
				}
			}
		}
	}

	return result
}

var coordinateHeatMap = make(map[Coordinate]int)

func aStar(rootNode *Node) *Node {

	// Keep track of a general list of nodes, and add list of children to priorityQueue
	var priorityQueue PriorityQueue
	var currentNode *Node
	var processedNodes []*Node
	coordinateHeatMap = make(map[Coordinate]int)

	// Generate all children of root node, and add children to priority queue
	rootNode.generateChildren()

	// Add children of root node to priority queue
	priorityQueue = rootNode.children

	flag:
	for len(priorityQueue) > 0 {

		// Check if user cancelled
		if stopSolver {
			return currentNode
		}

		// Sort priorityQueue
		sortByDistanceFrom(priorityQueue)

		// Generate random number between 1 and 10 such that 10% of the time it will pick a random node
		rand.Seed(time.Now().UnixNano())
		randNode := rand.Intn(11)

		// Pick random node or first node in priorityQueue
		if randNode == 5{
			currentNode = priorityQueue[rand.Intn(len(priorityQueue))]
		} else {
			currentNode = priorityQueue[0]
		}

		for _, processedNode := range processedNodes{
			if processedNode.currentPosition == currentNode.currentPosition {
				priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
				continue flag
			}
		}

		//printBoard(currentNode.board)

		// Update current board in browser
		if speedMode {
			updateBoard(currentNode.currentPosition, currentNode.parent.currentPosition)
		} else {
			fillGridFromBoard(currentNode.board)
		}

		// Check if current node is goal
		if currentNode.currentPosition == currentNode.goalPosition {
			if speedMode{
				fillGridFromBoard(currentNode.board)
			}
			return currentNode
		}

		// Update coordinateHeatmap
		updateCoordinateHeatMap(coordinateHeatMap, currentNode)

		// Generate children
		currentNode.generateChildren()
		nodesVisitedCounter++

		// Add children of current node to priority queue
		for _, cn := range currentNode.children{
			priorityQueue = append(priorityQueue, cn)
		}

		priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
		processedNodes = append(processedNodes, currentNode)

	}

	return &Node{}
}

func updateCoordinateHeatMap(coordinateHeatMap map[Coordinate]int, node *Node){
	coordinateHeatMap[node.currentPosition] += 1
}

func sortByDistanceFrom(nodes []*Node) {

	sort.Slice(nodes, func(i, j int) bool {

		// Calculate h values
		dix := float64(nodes[i].goalPosition.column - nodes[i].currentPosition.column)
		diy := float64(nodes[i].goalPosition.row - nodes[i].currentPosition.row)
		di :=  math.Sqrt(math.Pow(dix, 2) + math.Pow(diy, 2))

		djx := float64(nodes[j].goalPosition.column - nodes[j].currentPosition.column)
		djy := float64(nodes[j].goalPosition.row - nodes[j].currentPosition.row)
		dj :=  math.Sqrt(math.Pow(djx, 2) + math.Pow(djy, 2))

		// Calculate f values as f = g + h (g = heatmap value, h = pythagorean distance)
		fi := di + nodes[i].gValue
		fj := dj + nodes[j].gValue

		return fi < fj
	})
}

func aStarCancerous(rootNode *Node) *Node {

	// Keep track of a general list of nodes, and add list of children to priorityQueue
	var priorityQueue PriorityQueue
	var currentNode *Node
	var processedNodes []*Node

	// Generate all children of root node, and add children to priority queue
	rootNode.generateChildren()

	// Add children of root node to priority queue
	priorityQueue = rootNode.children

	flag:
	for len(priorityQueue) > 0 {

		// Check if user cancelled
		if stopSolver {
			return currentNode
		}

		// Sort priorityQueue
		sortByDistanceFromCancerous(priorityQueue)

		// Generate random number between 1 and 10 such that 10% of the time it will pick a random node
		rand.Seed(time.Now().UnixNano())
		randNode := rand.Intn(11)

		// Pick random node or first node in priorityQueue
		if randNode == 5{
			currentNode = priorityQueue[rand.Intn(len(priorityQueue))]
		} else {
			currentNode = priorityQueue[0]
		}

		for _, processedNode := range processedNodes{
			if processedNode.currentPosition == currentNode.currentPosition {
				priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
				continue flag
			}
		}

		//printBoard(currentNode.board)

		// Update current board in browser
		fillGridFromBoard(currentNode.board)

		// Check if current node is goal
		if currentNode.currentPosition == currentNode.goalPosition {
			return currentNode
		}

		// Generate children
		currentNode.generateChildren()
		nodesVisitedCounter++

		// Add children of current node to priority queue
		for _, cn := range currentNode.children{
			priorityQueue = append(priorityQueue, cn)
		}

		priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
		processedNodes = append(processedNodes, currentNode)
	}

	return &Node{}
}

func sortByDistanceFromCancerous(nodes []*Node) {
	sort.Slice(nodes, func(i, j int) bool {
		dix := float64(nodes[i].goalPosition.column - nodes[i].currentPosition.column)
		diy := float64(nodes[i].goalPosition.row - nodes[i].currentPosition.row)
		di :=  math.Sqrt(math.Pow(dix, 2) + math.Pow(diy, 2))

		djx := float64(nodes[j].goalPosition.column - nodes[j].currentPosition.column)
		djy := float64(nodes[j].goalPosition.row - nodes[j].currentPosition.row)
		dj :=  math.Sqrt(math.Pow(djx, 2) + math.Pow(djy, 2))
		return dj < di
	})
}
