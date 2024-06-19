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

// ShortestPathBFS 使用BFS查找最短路径
func (g *Graph) ShortestPathBFS(start, end int) []int {
	queue := []int{start}
	visited := make(map[int]bool)
	parent := make(map[int]int)

	found := false

	// BFS
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v == end {
			found = true
			break
		}

		for _, neighbor := range g.adj[v] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = v
				queue = append(queue, neighbor)
			}
		}
	}

	// 构建路径
	if found {
		path := []int{}
		for curr := end; curr != start; curr = parent[curr] {
			path = append([]int{curr}, path...)
		}
		path = append([]int{start}, path...)
		return path
	}

	return nil
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
	endVertex := 4

	fmt.Printf("Shortest path from %d to %d: ", startVertex, endVertex)
	shortestPath := graph.ShortestPathBFS(startVertex, endVertex)
	if shortestPath != nil {
		for i, node := range shortestPath {
			if i == len(shortestPath)-1 {
				fmt.Printf("%d\n", node)
			} else {
				fmt.Printf("%d -> ", node)
			}
		}
	} else {
		fmt.Println("No path found")
	}
}
