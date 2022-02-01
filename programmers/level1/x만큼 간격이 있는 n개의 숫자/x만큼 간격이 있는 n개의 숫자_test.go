package x만큼_간격이_있는_n개의_숫자

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(x int, n int) (result []int) {
	for i := 1; i <= n; i++ {
		result = append(result, x*i)
	}
	return
}

var x = []int{2, 4, -4}
var n = []int{5, 3, 2}
var answer = [][]int{{2, 4, 6, 8, 10}, {4, 8, 12}, {-4, -8}}

func TestSolution(t *testing.T) {
	for i := range answer {
		result := solution(x[i], n[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(12, 31)
	}
}
