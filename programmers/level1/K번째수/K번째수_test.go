package K번째수_test

import (
	"reflect"
	"sort"
	"testing"
)

func solution(array []int, commands [][]int) (ret []int) {
	for _, cmd := range commands {
		slice := append([]int{}, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(slice)
		ret = append(ret, slice[cmd[2]-1])
	}
	return
}

func TestSolution(t *testing.T) {
	testCase1 := solution([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}})
	answer1 := []int{5, 6, 3}
	if !reflect.DeepEqual(testCase1, answer1) {
		t.Error("testCase1 Error", testCase1, answer1)
	}

	testCase2 := solution([]int{2}, [][]int{{1, 1, 1}})
	answer2 := []int{2}
	if !reflect.DeepEqual(testCase2, answer2) {
		t.Error("testCase2 Error", testCase2, answer2)
	}
}
