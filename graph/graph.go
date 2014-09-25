package graph

import (
	"reflect"
)

type Node struct {
	Value     interface{}
	Neighbors []Node
}

func NewNode(value interface{}) Node {
	return Node{value, nil}
}

func FindShortestPath(start Node, end Node, path []Node) []Node {
	var shortest []Node
	path = append(path, start)

	if reflect.DeepEqual(start, end) {
		return path
	}

	shortest = nil

	for _, currentNode := range start.Neighbors {
		if nodeInPath(currentNode, path) {
			continue
		}
		newPath := FindShortestPath(currentNode, end, path)
		if newPath != nil && len(newPath) != 0 {
			if shortest == nil || (len(newPath) < len(shortest)) {
				shortest = newPath
			}
		}
	}
	return shortest
}

func nodeInPath(node Node, path []Node) bool {
	for _, currentNode := range path {
		if reflect.DeepEqual(node, currentNode) {
			return true
		}
	}
	return false
}
