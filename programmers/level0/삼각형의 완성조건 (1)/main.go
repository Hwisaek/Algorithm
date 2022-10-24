package main

func solution(sides []int) int {
	sum, max := 0, 0
	for _, side := range sides {
		if side > max {
			max = side
		}
		sum += side
	}
	if sum > max*2 {
		return 1
	}
	return 2
}
