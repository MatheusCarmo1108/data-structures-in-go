// About the Graph implemented -  directed graph with representation as adjacency list
package main

import "fmt"

// Graph structure - represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex structure - represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex - adds a Vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Add Edge - adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil { // if one of them are a empty Vertex
		err := fmt.Errorf("WARNING: invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) { // if the Edge already exists
		err := fmt.Errorf("WARNING: existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edege
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print - Will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, myVertices := range g.vertices {
		fmt.Printf("\nVertex %v : ", myVertices.key)
		for _, adjacentVertices := range myVertices.adjacent {
			fmt.Printf(" %v ", adjacentVertices.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)

	test.Print()
}
