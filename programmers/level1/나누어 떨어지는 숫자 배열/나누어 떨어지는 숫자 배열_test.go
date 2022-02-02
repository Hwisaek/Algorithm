package 나누어_떨어지는_숫자_배열

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func solution(arr []int, divisor int) (result []int) {
	for _, v := range arr {
		if v%divisor == 0 {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		result = []int{-1}
	} else {
		sort.Ints(result)
	}
	return
}

func TestSolution(t *testing.T) {
	arr := [][]int{{5, 9, 7, 10}, {2, 36, 1, 3}, {3, 2, 6}}
	divisor := []int{5, 1, 10}
	answer := [][]int{{5, 10}, {1, 2, 3, 36}, {-1}}

	for i := range answer {
		result := solution(arr[i], divisor[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2)
	}
}
