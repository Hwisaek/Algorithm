package main

func solution(numbers []int, direction string) []int {
	switch direction {
	case "left":
		return append(numbers[1:], numbers[0])
	default:
		return append([]int{numbers[len(numbers)-1]}, numbers[:len(numbers)-1]...)
	}
}
