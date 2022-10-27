package main

import "math"

func solution(n int) int {
	sqrt := math.Sqrt(float64(n))
	if sqrt-float64(int(sqrt)) > 0 {
		return 2
	}
	return 1
}
