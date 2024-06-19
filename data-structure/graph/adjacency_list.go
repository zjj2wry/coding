package main

import "fmt"

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
	// 无向图，因此双向都需要添加
	g.adj[v1] = append(g.adj[v1], v2)
	g.adj[v2] = append(g.adj[v2], v1)
}

// PrintGraph 打印图的邻接表表示
func (g *Graph) PrintGraph() {
	for v, neighbors := range g.adj {
		fmt.Printf("Vertex %d -> ", v)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()

	// 添加顶点
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	// 添加边
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	// 打印图的邻接表表示
	graph.PrintGraph()
}
