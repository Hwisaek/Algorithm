package main

import (
	"fmt"
)

func solution(array []int) (result int) {
	result = -1
	count := 0

	m := make(map[int]int)
	for _, n := range array {
		m[n]++
	}

	for k, v := range m {
		if v > count {
			result = k
			count = v
		} else if v == count {
			result = -1
		}
	}
	return
}

func main() {
	fmt.Println(solution([]int{1, 1, 2, 2, 3}))
	fmt.Println(solution([]int{1, 1, 2, 2, 3, 1}))
}
