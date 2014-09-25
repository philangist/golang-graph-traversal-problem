package main

import (
	"fmt"

	"github.com/philangist/golang-graph-traversal-problem/graph"
)

func main() {
	A := graph.NewNode("A")
	B := graph.NewNode("B")
	C := graph.NewNode("C")
	D := graph.NewNode("D")
	E := graph.NewNode("E")
	F := graph.NewNode("F")

	A.Neighbors = append(A.Neighbors, B)
	A.Neighbors = append(A.Neighbors, C)

	B.Neighbors = append(B.Neighbors, D)
	B.Neighbors = append(B.Neighbors, E)

	C.Neighbors = append(C.Neighbors, F)

	fmt.Println("A.Neighbors is", D.Neighbors)
	fmt.Println("B.Neighbors is", B.Neighbors)
	fmt.Println("C.Neighbors is", C.Neighbors)
	fmt.Println("The shortest path from A to D is",
		graph.FindShortestPath(A, F, []graph.Node{}))

}
