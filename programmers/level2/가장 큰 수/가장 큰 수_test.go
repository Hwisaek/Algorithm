package 가장_큰_수_test

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func solution(numbers []int) (result string) {
	strNumbers := make([]string, len(numbers))
	for i, number := range numbers {
		str := strconv.Itoa(number)
		strNumbers[i] = str
	}

	sort.Slice(strNumbers, func(i, j int) bool {
		return strNumbers[i]+strNumbers[i]+strNumbers[i] > strNumbers[j]+strNumbers[j]+strNumbers[j]
	})

	result = strings.Join(strNumbers, "")
	preResult, _ := strconv.Atoi(result)
	if preResult == 0 {
		result = "0"
	}
	return
}

func TestSolution(t *testing.T) {
	testCase := [][]int{{6, 10, 2}, {3, 30, 34, 5, 9}, {0, 0, 0}, {9, 998}, {9, 1000}}
	answer := []string{"6210", "9534330", "0", "9998", "91000"}

	for i := range testCase {
		result := solution(testCase[i])
		if result != answer[i] {
			t.Error(fmt.Sprintf("TestCase%d error. TestCase: %v, answer: %v", i, result, answer[i]))
		}
	}
}
