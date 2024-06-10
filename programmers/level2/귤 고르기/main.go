package main

import "sort"

func solution(k int, tangerine []int) int {
	sum := k
	count := map[int]int{}
	for _, v := range tangerine {
		count[v]++
	}

	slice := make([]int, 0)
	for _, v := range count {
		slice = append(slice, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	for i, e := range slice {
		sum -= e
		if sum <= 0 {
			return i + 1
		}
	}

	return k
}

func main() {
	solution(6, []int{1, 3, 2, 5, 4, 5, 2, 3})
}
