package 문자열_내림차순으로_배치하기_test

import (
	"fmt"
	"sort"
	"testing"
)

func solution(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})
	return string(bytes)
}

func TestSolution(t *testing.T) {
	numbers := []string{"Zbcdefg"}
	answer := []string{"gfedcbZ"}

	for i := range answer {
		result := solution(numbers[i])
		if result != answer[i] {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
