package main

import (
	"fmt"
)

func solution(num_list []int) []int {
	for i := len(num_list) - 1; i >= 0; i-- {
		num_list = append(num_list, num_list[i])
	}
	return num_list[len(num_list)/2:]
}

func main() {
	fmt.Println(solution([]int{1, 0, 1, 1, 1, 3, 5}))
}
