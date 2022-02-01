package 가운데_글자_가져오기_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(s string) string {
	l := len(s)
	n := l / 2
	if l%2 == 0 {
		return s[n-1 : n+1]
	}
	return s[n : n+1]
}

func TestSolution(t *testing.T) {
	s := []string{"abcde", "qwer"}
	answer := []string{"c", "we"}

	for i := range answer {
		result := solution(s[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
