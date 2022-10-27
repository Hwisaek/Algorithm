package 중앙값_구하기

import "sort"

func solution(array []int) int {
	sort.Ints(array)
	return array[len(array)/2]
}
