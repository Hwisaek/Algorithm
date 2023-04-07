package main

import (
	"sort"
)

func solution(distance int, rocks []int, n int) int {
	sort.Ints(rocks)
	left := 1
	right := distance
	answer := 0

	for left <= right {
		mid := (left + right) / 2
		prev := 0
		removeCnt := 0

		for i := 0; i < len(rocks); i++ {
			if rocks[i]-prev < mid {
				removeCnt++
			} else {
				prev = rocks[i]
			}
		}

		if distance-prev < mid {
			removeCnt++
		}

		if removeCnt <= n {
			left = mid + 1
			answer = mid
		} else {
			right = mid - 1
		}
	}

	return answer
}
