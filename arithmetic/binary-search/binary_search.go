package main

import (
	"fmt"
)

func main() {
	value := 5
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(value, s)
	fmt.Println(result)

	result = binarySearch(11, s)
	fmt.Println(result)
}

func binarySearch(value int, s []int) int {
	i := 0
	j := len(s) - 1

	for i <= j {
		mid := (i + j) / 2
		if s[mid] == value {
			return mid
		} else if s[mid] > value {
			j = mid - 1
		} else if s[mid] < value {
			i = mid + 1
		}
	}
	return -1
}
