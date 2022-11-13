package main

import "fmt"

func solution(numbers []int, k int) (answer int) {
	answer = 1
	length := len(numbers)
	for k > 1 {
		fmt.Println(answer)
		answer += 2
		if answer > length {
			answer = answer%length + 1
		}
		k--
	}
	return
}

func main() {
	solution([]int{1, 2, 3, 4}, 2)
}
