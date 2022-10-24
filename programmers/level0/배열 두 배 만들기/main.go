package main

func solution(numbers []int) (result []int) {
	result = make([]int, len(numbers))
	for i, number := range numbers {
		result[i] = number * 2
	}
	return
}
