package 프린터

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(priorities []int, location int) (count int) {
	for len(priorities) > 0 {
		J := priorities[0]
		priorities = priorities[1:]
		location--
		if isPrior(J, priorities) {
			count++
			if location == -1 {
				return count
			}
		} else {
			priorities = append(priorities, J)
			if location < 0 {
				location = len(priorities) - 1
			}
		}
	}
	return
}

func isPrior(j int, priorities []int) bool {
	max := 0
	for _, priority := range priorities {
		if priority > max {
			max = priority
		}
	}
	return j >= max
}

func TestSolution(t *testing.T) {
	priorities := [][]int{{2, 1, 3, 2}, {1, 1, 9, 1, 1, 1}}
	location := []int{2, 0}
	answer := []int{1, 5}

	for i := range answer {
		result := solution(priorities[i], location[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Testcase %v Failed. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	priorities := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		priorities = append(priorities, i)
	}
	for i := 0; i < b.N; i++ {
		solution(priorities, 99)
	}
}
