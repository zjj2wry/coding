package main

import "fmt"

// Graph 结构体表示图
type Graph struct {
	vertices int      // 图的顶点数
	matrix   [][]bool // 邻接矩阵
}

// NewGraph 创建一个新的图实例
func NewGraph(numVertices int) *Graph {
	matrix := make([][]bool, numVertices)
	for i := range matrix {
		matrix[i] = make([]bool, numVertices)
	}
	return &Graph{
		vertices: numVertices,
		matrix:   matrix,
	}
}

// AddEdge 添加一条边到图中
func (g *Graph) AddEdge(v1, v2 int) {
	g.matrix[v1][v2] = true
	g.matrix[v2][v1] = true // 无向图，双向添加
}

// PrintGraph 打印图的邻接矩阵表示
func (g *Graph) PrintGraph() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if g.matrix[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph(5)

	// 添加边
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	// 打印图的邻接矩阵表示
	graph.PrintGraph()
}
