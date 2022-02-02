package 수박수박수박수박수박수

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func solution(n int) string {
	var result strings.Builder

	for i := range make([]int, n) {
		if i%2 == 0 {
			result.WriteString("수")
		} else {
			result.WriteString("박")
		}
	}

	return result.String()
}

func TestSolution(t *testing.T) {
	n := []int{3, 4}
	answer := []string{"수박수", "수박수박"}

	for i := range answer {
		result := solution(n[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(10000)
	}
}
