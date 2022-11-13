package main

import "math"

func solution(array []int, n int) (answer int) {
	answer = 0
	dist := 100
	for _, e := range array {
		diff := int(math.Abs(float64(e - n)))
		if dist > diff {
			answer = e
			dist = diff
		} else if dist == diff && answer > e {
			answer = e
		}
	}
	return
}
