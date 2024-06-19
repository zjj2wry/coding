package main

import (
	"fmt"
)

// Graph 结构体表示图
type Graph struct {
	vertices int           // 图的顶点数
	adj      map[int][]int // 邻接表
}

// NewGraph 创建一个新的图实例
func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}
}

// AddVertex 添加一个顶点到图中
func (g *Graph) AddVertex(v int) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = []int{}
		g.vertices++
	}
}

// AddEdge 添加一条边到图中
func (g *Graph) AddEdge(v1, v2 int) {
	g.adj[v1] = append(g.adj[v1], v2)
	g.adj[v2] = append(g.adj[v2], v1) // 无向图，双向添加
}

// BFS 执行广度优先搜索
func (g *Graph) BFS(startVertex int) {
	visited := make(map[int]bool)
	queue := []int{startVertex}

	visited[startVertex] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", v)

		for _, neighbor := range g.adj[v] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
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
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1)

	startVertex := 0
	fmt.Printf("BFS traversal starting from vertex %d: ", startVertex)
	graph.BFS(startVertex)
	fmt.Println()
}
