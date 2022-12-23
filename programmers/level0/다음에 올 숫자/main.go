package main

import "fmt"

func solution(common []int) int {
	if common[0] != 0 {
		y := common[1] / common[0]
		if common[2] == common[1]*y {
			return common[len(common)-1] * y
		}
	}
	x := common[1] - common[0]
	return common[len(common)-1] + x
}

func main() {
	fmt.Println(solution([]int{2, 4, 8}))
	fmt.Println(solution([]int{1, 2, 3, 4}))
	fmt.Println(solution([]int{0, 0, 0}))
	fmt.Println(solution([]int{0, -1}))
}
