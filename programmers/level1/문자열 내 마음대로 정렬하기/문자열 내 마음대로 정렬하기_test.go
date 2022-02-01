package 문자열_내_마음대로_정렬하기_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func solution(strings []string, n int) []string {
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] != strings[j][n] {
			return strings[i][n] < strings[j][n]
		}
		return strings[i] < strings[j]

	})
	return strings
}

func TestSolution(t *testing.T) {
	strings := [][]string{{"sun", "bed", "car"}, {"abce", "abcd", "cdx"}}
	n := []int{1, 2}
	answer := [][]string{{"car", "bed", "sun"}, {"abcd", "abce", "cdx"}}

	for i := range answer {
		result := solution(strings[i], n[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
