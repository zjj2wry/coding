package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 5, 34, 2, 6, 8, 9, 0, 0, 5}
	s2 := s1
	s3 := s1
	s4 := s1
	Bubble(s1)
	Select(s2)
	Insert(s3)
	Insert2(s4)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

// 相邻之间互相比较，如果前面的比后面的小，就交换位置
func Bubble(s []int) {
	if len(s) <= 1 {
		return
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				swap(s, i, j)
			}
		}
	}
}

// 每次找到最小值所在的坐标然后交换值, 选择排序是一种不稳定的排序
func Select(s []int) {
	if len(s) <= 1 {
		return
	}
	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if i != min {
			swap(s, min, i)
		}
	}
}

// 往一个已经排序的数组中插入新的数据
func Insert(s []int) {
	if len(s) <= 1 {
		return
	}

	for i := 1; i < len(s); i++ {
		j := i
		// 如果后面一个数比前面的数小，就交换
		for j > 0 && s[j] < s[j-1] {
			swap2(s, j, j-1)
			j--
		}
	}
}

// 和 insert1 相比，减少了 swap 的次数，插入排序的平均速度大于冒泡和选择
func Insert2(s []int) {
	if len(s) <= 1 {
		return
	}
	var j int
	for i := 1; i < len(s); i++ {
		temp := s[i]
		// 如果后面一个数比前面的数小，就交换
		for j = i; j > 0 && temp < s[j-1]; j-- {
			s[j] = s[j-1]
		}
		s[j] = temp
	}
}

func swap(s []int, i, j int) {
	temp := s[j]
	s[j] = s[i]
	s[i] = temp
}

// 使用与或交换，不用临时的变量
func swap2(s []int, i, j int) {
	s[i] ^= s[j]
	s[j] ^= s[i]
	s[i] ^= s[j]
}
