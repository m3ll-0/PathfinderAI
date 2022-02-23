package main

type PriorityQueue []*Node

type SolverMethod func(node *Node) *Node

type UpdateBoardMethod func(prevPlayerCoordinate Coordinate, curPlayerCoordinate Coordinate)
type GenerateBoardMethod func(board Board)
type FillBoardMethod func(board Board)

type Coordinate struct {
	row int
	column int
}