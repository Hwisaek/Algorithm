package 모의고사

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(answers []int) (result []int) {
	checked := [][]int{{1, 2, 3, 4, 5}, {2, 1, 2, 3, 2, 4, 2, 5}, {3, 3, 1, 1, 2, 2, 4, 4, 5, 5}}
	scores := []int{0, 0, 0}
	for i, answer := range answers {
		for j := range scores {
			if checked[j][i%len(checked[j])] == answer {
				scores[j]++
			}
		}
	}
	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	for i := range scores {
		if scores[i] == max {
			result = append(result, i+1)
		}
	}
	return
}

func TestSolution(t *testing.T) {
	answers := [][]int{{1, 2, 3, 4, 5}, {1, 3, 2, 4, 2}}
	answer := [][]int{{1}, {1, 2, 3}}

	for i := range answer {
		result := solution(answers[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Testcase %v Failed. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	answers := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		answers = append(answers, i)
	}
	for i := 0; i < b.N; i++ {
		solution(answers)
	}
}
