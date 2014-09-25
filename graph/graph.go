package graph

import "fmt"

type Node struct {
	Value interface{}
  Neighbors []Node
}

func NewNode(value interface{}) Node {
  return Node{value, nil}
}

func (this Node) AddNeighbors(neighbors []Node) {
  neighbors_copy := this.Neighbors
  for _, node := range neighbors {
    fmt.Println("node is", node.Value)
    neighbors_copy = append(neighbors_copy, node)
  }
  this.Neighbors = neighbors_copy
}
