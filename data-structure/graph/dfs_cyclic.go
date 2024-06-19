package main

import "fmt"

type Graph struct {
	vertices int
	adj      map[int][]int
	visited  []bool
	recStack []bool
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adj:      make(map[int][]int),
		visited:  make([]bool, vertices),
		recStack: make([]bool, vertices),
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.adj[v1] = append(g.adj[v1], v2)
}

func (g *Graph) isCyclicUtil(v int) bool {
	if !g.visited[v] {
		g.visited[v] = true
		g.recStack[v] = true

		for _, neighbor := range g.adj[v] {
			if !g.visited[neighbor] && g.isCyclicUtil(neighbor) {
				return true
			} else if g.recStack[neighbor] {
				return true
			}
		}
	}

	g.recStack[v] = false
	return false
}

func (g *Graph) IsCyclic() bool {
	for i := 0; i < g.vertices; i++ {
		if g.isCyclicUtil(i) {
			return true
		}
	}
	return false
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	if g.IsCyclic() {
		fmt.Println("Graph contains cycle")
	} else {
		fmt.Println("Graph doesn't contain cycle")
	}
}
