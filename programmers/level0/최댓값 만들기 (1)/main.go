package 최댓값_만들기__1_

import "sort"

func solution(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-1] * numbers[len(numbers)-2]
}
