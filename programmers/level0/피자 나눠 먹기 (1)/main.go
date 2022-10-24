package main

import "math"

func solution(n int) int {
	return int(math.Ceil(float64(n) / 7.))
}
