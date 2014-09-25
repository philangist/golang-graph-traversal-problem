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

  A.Neighbors = append(A.Neighbors, B)
  A.Neighbors = append(A.Neighbors, C)

  B.Neighbors = append(B.Neighbors, C)
  B.Neighbors = append(B.Neighbors, D)

  C.Neighbors = append(C.Neighbors, D)
  D.Neighbors = append(D.Neighbors, C)
  E.Neighbors = append(E.Neighbors, F)
  F.Neighbors = append(F.Neighbors, C)

  fmt.Println("A.Neighbors is", A.Neighbors)
  fmt.Println("B.Neighbors is", B.Neighbors)
  fmt.Println("C.Neighbors is", C.Neighbors)
  fmt.Println("D.Neighbors is", D.Neighbors)
  fmt.Println("E.Neighbors is", E.Neighbors)
  fmt.Println("F.Neighbors is", F.Neighbors)

}
