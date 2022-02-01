package 없는_숫자_더하기_test

import (
	"fmt"
	"testing"
)

func solution(numbers []int) (result int) {
	result = 45
	for _, number := range numbers {
		result -= number
	}
	return
}

func TestSolution(t *testing.T) {
	numbers := [][]int{{1, 2, 3, 4, 6, 7, 8, 0}, {5, 8, 4, 0, 6, 7, 9}}
	answer := []int{14, 6}

	for i := range answer {
		result := solution(numbers[i])
		if result != answer[i] {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
