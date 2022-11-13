package main

import "sort"

func solution(emergency []int) (result []int) {
	arr := [][]int{}
	for i, e := range emergency {
		arr = append(arr, []int{i, e})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	result = make([]int, len(arr))
	for i, e := range arr {
		result[e[0]] = i + 1
	}
	return
}
