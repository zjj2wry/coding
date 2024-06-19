package main

import (
	"fmt"
)

// Graph 结构体表示图
type Graph struct {
	vertices int           // 图的顶点数
	adj      map[int][]int // 邻接表
	visited  map[int]bool  // 记录节点是否被访问过
}

// NewGraph 创建一个新的图实例
func NewGraph() *Graph {
	return &Graph{
		adj:     make(map[int][]int),
		visited: make(map[int]bool),
	}
}

// AddVertex 添加一个顶点到图中
func (g *Graph) AddVertex(v int) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = []int{}
		g.visited[v] = false
		g.vertices++
	}
}

// AddEdge 添加一条边到图中
func (g *Graph) AddEdge(v1, v2 int) {
	g.adj[v1] = append(g.adj[v1], v2)
	g.adj[v2] = append(g.adj[v2], v1) // 无向图，双向添加
}

// DFS 执行深度优先搜索
func (g *Graph) DFS(startVertex int) {
	g.visited[startVertex] = true
	fmt.Printf("%d ", startVertex)

	for _, neighbor := range g.adj[startVertex] {
		if !g.visited[neighbor] {
			g.DFS(neighbor)
		}
	}
}

func main() {
	graph := NewGraph()

	// 添加顶点
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	// 添加边
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	startVertex := 2
	fmt.Printf("DFS traversal starting from vertex %d: ", startVertex)
	graph.DFS(startVertex)
	fmt.Println()
}
