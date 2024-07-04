package main

import (
	"fmt"
	"math"
	"sync"
)

// Vector 表示一个向量
type Vector []float64

// VectorDB 表示一个简单的向量数据库
type VectorDB struct {
	vectors map[int]Vector
	mutex   sync.Mutex
}

// NewVectorDB 创建一个新的向量数据库
func NewVectorDB() *VectorDB {
	return &VectorDB{
		vectors: make(map[int]Vector),
	}
}

// AddVector 向数据库中添加一个向量
func (db *VectorDB) AddVector(id int, vector Vector) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.vectors[id] = vector
}

// GetVector 从数据库中获取一个向量
func (db *VectorDB) GetVector(id int) (Vector, bool) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	vector, found := db.vectors[id]
	return vector, found
}

// EuclideanDistance 计算两个向量之间的欧氏距离
func EuclideanDistance(v1, v2 Vector) float64 {
	if len(v1) != len(v2) {
		return math.Inf(1)
	}
	sum := 0.0
	for i := range v1 {
		diff := v1[i] - v2[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// FindNearest 找到数据库中距离目标向量最近的向量
func (db *VectorDB) FindNearest(target Vector) (int, Vector, float64) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	var nearestID int
	var nearestVector Vector
	minDistance := math.Inf(1)

	for id, vector := range db.vectors {
		distance := EuclideanDistance(vector, target)
		if distance < minDistance {
			minDistance = distance
			nearestID = id
			nearestVector = vector
		}
	}

	return nearestID, nearestVector, minDistance
}

func main() {
	db := NewVectorDB()

	// 添加一些向量
	db.AddVector(1, Vector{1.0, 2.0, 3.0})
	db.AddVector(2, Vector{4.0, 5.0, 6.0})
	db.AddVector(3, Vector{7.0, 8.0, 9.0})

	// 获取一个向量
	vector, found := db.GetVector(2)
	if found {
		fmt.Printf("Vector with ID 2: %v\n", vector)
	} else {
		fmt.Println("Vector with ID 2 not found")
	}

	// 查找最近的向量
	target := Vector{5.0, 5.0, 5.0}
	nearestID, nearestVector, distance := db.FindNearest(target)
	fmt.Printf("Nearest vector to %v is ID %d with vector %v and distance %f\n", target, nearestID, nearestVector, distance)
}
