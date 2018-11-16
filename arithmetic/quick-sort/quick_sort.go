	fmt.Println(s1)
	fmt.Println(s1)
	fmt.Println(s2)
}

func MergeSort(s []int) {
	sort(s, 0, len(s)-1)
}

func sort(s []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(s, left, right)
	sort(s, left, index-1)
	sort(s, index+1, right)
}

// 找到一个枢纽，不断的交换左右的位置
func partition(s []int, left, right int) int {
	pivotPtr := right
	pivot := s[right]
	right = right - 1

	for {
		for s[left] < pivot {
			left++
		}
		// avoid len(s)==1, right will == 0
		for right > 0 && s[right] > pivot {
			right--
		}
		if left >= right {
			fmt.Println(left, right)
			break
		}

		swap(s, left, right)
	}

	swap(s, left, pivotPtr)
	return left
}

func swap(s []int, i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
