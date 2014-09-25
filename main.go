package main

import (
  "fmt"

  "github.com/philangist/golang-graph-traversal-problem/graph"
)


func main (){
  A := graph.NewNode("A")
  B := graph.NewNode("B")
  C := graph.NewNode("C")
  D := graph.NewNode("D")
  E := graph.NewNode("E")
  F := graph.NewNode("F")

  A.AddNeighbors([]graph.Node{B, C})
  B.AddNeighbors([]graph.Node{C, D})
  C.AddNeighbors([]graph.Node{D})
  D.AddNeighbors([]graph.Node{C})
  E.AddNeighbors([]graph.Node{F})
  F.AddNeighbors([]graph.Node{C})

  fmt.Println(A.Value)

}
