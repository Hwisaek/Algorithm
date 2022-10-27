package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	if !reflect.DeepEqual(solution(1, 2, 3, 4), []int{5, 4}) {
		t.Error("error")
	}
	if !reflect.DeepEqual(solution(9, 2, 1, 3), []int{29, 6}) {
		t.Error("error")
	}
}
