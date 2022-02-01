package 두_쩡수_사이의_합_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(a int, b int) (result int) {
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		result += i
	}
	return
}

func TestSolution(t *testing.T) {
	a := []int{3, 3, 5}
	b := []int{5, 3, 3}
	answer := []int{12, 3, 12}

	for i := range a {
		result := solution(a[i], b[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("TestCase %d: Error. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
