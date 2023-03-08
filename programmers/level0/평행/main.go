package main

import "math"

func slope(p1, p2 []int) float64 {
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func isParallel(p1, p2, p3, p4 []int) bool {
	slope1 := slope(p1, p2)
	slope2 := slope(p3, p4)

	return math.Abs(slope1) == math.Abs(slope2)
}

func solution(dots [][]int) int {
	p1 := dots[0]
	p2 := dots[1]
	p3 := dots[2]
	p4 := dots[3]

	if isParallel(p1, p2, p3, p4) {
		return 1
	} else {
		return 0
	}
}

func main() {
	solution([][]int{{3, 5}, {4, 1}, {2, 4}, {5, 10}})
}
