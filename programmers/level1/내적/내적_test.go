package 내적_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(a []int, b []int) (result int) {
	for i, v := range a {
		result += v * b[i]
	}
	return
}

func TestSolution(t *testing.T) {
	a := [][]int{{1, 2, 3, 4}, {-1, 0, 1}}
	b := [][]int{{-3, -1, 0, 2}, {1, 0, -1}}
	answer := []int{3, -2}

	for i := range a {
		result := solution(a[i], b[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("TestCase %d: Error. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
