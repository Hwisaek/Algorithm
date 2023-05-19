package main

import "sort"

func solution(k int, tangerine []int) int {
	count := map[int]int{}
	for _, v := range tangerine {
		count[v]++
	}

	mapToSlice := make([]int, 0)
	for _, v := range count {
		mapToSlice = append(mapToSlice, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mapToSlice)))

	return 0
}
