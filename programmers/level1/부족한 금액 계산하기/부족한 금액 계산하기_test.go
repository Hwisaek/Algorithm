package 부족한_금액_계산하기_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(price int, money int, count int) (result int) {
	sum := 0
	for i := 1; i <= count; i++ {
		sum += price * i
	}
	result = sum - money
	if result < 0 {
		return 0
	}
	return
}

func TestSolution(t *testing.T) {
	price := []int{3}
	money := []int{20}
	count := []int{4}
	answer := []int{10}

	for i := range answer {
		result := solution(price[i], money[i], count[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Error %v. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
