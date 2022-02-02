package 서울에서_김서방_찾기

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(seoul []string) (result string) {
	for i, s := range seoul {
		if s == "Kim" {
			result = fmt.Sprintf("김서방은 %v에 있다", i)
		}
	}
	return
}

func TestSolution(t *testing.T) {
	seoul := [][]string{{"Jane", "Kim"}}
	answer := []string{"김서방은 1에 있다"}

	for i := range answer {
		result := solution(seoul[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]string{"Jane", "Kim"})
	}
}
