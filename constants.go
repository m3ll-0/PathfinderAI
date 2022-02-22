package main

type PriorityQueue []*Node

type SolverMethod func(node *Node) *Node

type Coordinate struct {
	row int
	column int
}