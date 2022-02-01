package 약수의_개수와_덧셈_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(left int, right int) (result int) {
	for i := left; i <= right; i++ {
		if getNumOfDivisors(i)%2 == 0 {
			result += i
		} else {
			result -= i
		}
	}
	return
}

func getNumOfDivisors(n int) (count int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return
}

var left = []int{13, 24}
var right = []int{17, 27}
var answer = []int{43, 52}

func TestSolution(t *testing.T) {
	for i := range answer {
		result := solution(left[i], right[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(1, 1000)
	}
}
