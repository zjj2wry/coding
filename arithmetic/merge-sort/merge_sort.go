package main

import "fmt"

func main() {
	e1 := []int{1, 2, 3, 4, 3, 4, 5, 6}
	e2 := []int{11, 233, 1213, 433, 23, 454, 52, 234}
	MergeSort(e1)
	MergeSort(e)

	fmt.Println(e1)
	fmt.Println(e2)
}

func MergeSort(s []int) {
	sort(s, 0, len(s)-1)
}

func sort(s []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		sort(s, left, mid)
		sort(s, mid+1, right)
		merge(s, left, mid, right)
	}
}

// 对两个有序的数组进行排序
func merge(s []int, left, mid, right int) {
	k := 0
	i := left
	j := mid + 1
	temp := make([]int, len(s))
	for i <= mid && j <= right {
		if s[i] <= s[j] {
			temp[k] = s[i]
			k++
			i++
		} else {
			temp[k] = s[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = s[i]
		k++
		i++
	}

	for j <= right {
		temp[k] = s[j]
		k++
		j++
	}

	k = 0
	for left < right {
		s[left] = temp[k]
		left++
		k++
	}
}
