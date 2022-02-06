package 다리를_지나는_트럭

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func solution(citations []int) (result int) {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, citation := range citations {
		if citation < i+1 {
			break
		} else {
			result = i + 1
		}
	}
	return
}

func TestSolution(t *testing.T) {
	citations := [][]int{{3, 0, 6, 1, 5}, {7, 7, 8, 9}}
	answer := []int{3, 4}

	for i := range answer {
		result := solution(citations[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Testcase %v Failed. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	citations := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		citations = append(citations, i)
	}
	for i := 0; i < b.N; i++ {
		solution(citations)
	}
}
