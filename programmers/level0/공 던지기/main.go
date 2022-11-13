package main

func solution(numbers []int, k int) (answer int) {
	answer = 1
	length := len(numbers)
	for k > 1 {
		answer += 2
		if answer > length {
			answer = answer % length
		}
		k--
	}
	return
}

func main() {
	solution([]int{1, 2, 3, 4}, 2)
	solution([]int{1, 2, 3, 4, 5, 6}, 5)
	solution([]int{1, 2, 3}, 3)
}
