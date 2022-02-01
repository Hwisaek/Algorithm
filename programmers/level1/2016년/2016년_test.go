package _2016년_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(a int, b int) (result string) {
	dates := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	days := []string{"FRI", "SAT", "SUN", "MON", "TUE", "WED", "THU"}

	date := -1
	for i := 0; i < a-1; i++ {
		date += dates[i]
	}
	result = days[(date+b)%7]
	return
}

var a = []int{5, 1, 1, 1, 1, 1, 1, 1, 1, 12}
var b = []int{24, 1, 3, 2, 4, 5, 6, 7, 8, 31}
var answer = []string{"TUE", "FRI", "SUN", "SAT", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

func TestSolution(t *testing.T) {
	for i := range answer {
		result := solution(a[i], b[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(12, 31)
	}
}
