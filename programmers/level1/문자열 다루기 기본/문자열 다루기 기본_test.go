package 문자열_다루기_기본_test

import (
	"fmt"
	"strconv"
	"testing"
)

func solution(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}
	_, err := strconv.Atoi(s)
	if err == nil {
		return false
	}
	return true
}

func TestSolution(t *testing.T) {
	s := []string{"a234", "1234"}
	answer := []bool{false, true}

	for i := range answer {
		result := solution(s[i])
		if result != answer[i] {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
