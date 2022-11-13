package main

import "sort"

func solution(numbers []int) int {
	l := len(numbers)
	sort.Ints(numbers)
	min, max := numbers[0]*numbers[1], numbers[l-1]*numbers[l-2]
	if min > max {
		return min
	}
	return max
}
