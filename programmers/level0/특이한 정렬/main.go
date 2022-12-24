package main

import (
	"math"
	"sort"
)

func solution(numlist []int, n int) []int {
	sort.Slice(numlist, func(i, j int) bool {
		if math.Abs(float64(n-numlist[i])) == math.Abs(float64(n-numlist[j])) {
			return numlist[i] > numlist[j]
		} else {
			return math.Abs(float64(n-numlist[i])) < math.Abs(float64(n-numlist[j]))
		}
	})
	return numlist
}
