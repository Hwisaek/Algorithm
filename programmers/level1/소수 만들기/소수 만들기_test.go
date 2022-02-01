package 문자열_다루기_기본_test

import (
	"fmt"
	"math"
	"testing"
)

func solution(nums []int) (answer int) {
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			for k := j + 1; k < lenNums; k++ {
				if isPrime(nums[i] + nums[j] + nums[k]) {
					answer++
				}
			}
		}
	}
	return answer
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func TestSolution(t *testing.T) {
	s := [][]int{{1, 2, 3, 4}, {1, 2, 7, 6, 4}}
	answer := []int{1, 4}

	for i := range answer {
		result := solution(s[i])
		if result != answer[i] {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
