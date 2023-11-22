package main

import (
	"fmt"
)

func solution(rank []int, attendance []bool) int {
	l := len(rank) + 1
	a, b, c := l, l, l
	var ai, bi, ci int

	for i, v := range rank {
		if !attendance[i] {
			continue
		}
		if v < c {
			c = v
			ci = i
			if v < b {
				c, b = b, c
				ci, bi = bi, ci
				if v < a {
					a, b = b, a
					ai, bi = bi, ai
				}
			}
		}
	}

	return 10000*ai + 100*bi + ci
}

func main() {
	//fmt.Println(solution([]int{3, 7, 2, 5, 4, 6, 1}, []bool{false, true, true, true, true, false, false}))
	fmt.Println(solution([]int{1, 2, 3}, []bool{true, true, true}))
}
