package 기능개발_test

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(progresses []int, speeds []int) (realAnswer []int) {
	var answer []int

	lenProgresses := len(progresses)
	for i := 0; i < lenProgresses; i++ {
		t := 0
		for {
			progress := progresses[i] + speeds[i]*t
			t++
			if progress >= 100 {
				break
			}
		}
		answer = append(answer, t)
	}

	for {
		if len(answer) == 0 {
			break
		}
		val := answer[0]
		t := 0
		lenAnswer := len(answer)
		for i := 0; i < lenAnswer; i++ {
			if answer[0] > val {
				break
			}
			answer = answer[1:]
			t++
		}
		realAnswer = append(realAnswer, t)
	}
	return
}

func TestSolution(t *testing.T) {
	progresses := [][]int{{93, 30, 55}, {95, 90, 99, 99, 80, 99}}
	speeds := [][]int{{1, 30, 5}, {1, 1, 1, 1, 1, 1}}
	answer := [][]int{{2, 1}, {1, 3, 2}}

	for i := range progresses {
		result := solution(progresses[i], speeds[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("TestCase %d: Error. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}
