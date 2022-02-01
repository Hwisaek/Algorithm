package 문자열_내_마음대로_정렬하기_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(absolutes []int, signs []bool) (sum int) {
	for i, absolute := range absolutes {
		if signs[i] {
			sum += absolute
		} else {
			sum -= absolute
		}
	}
	return
}

func TestSolution(t *testing.T) {
	absolutes := [][]int{{4, 7, 12}, {1, 2, 3}}
	signs := [][]bool{{true, false, true}, {false, false, true}}
	answer := []int{9, 0}

	for i := range answer {
		result := solution(absolutes[i], signs[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
