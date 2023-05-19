package main

import "sort"

func solution(k int, m int, score []int) (price int) {
	sort.Slice(score, func(i, j int) bool { return score[i] > score[j] })
	for i := 0; i+m <= len(score); i += m {
		price += m * score[i+m-1]
	}
	return
}

func main() {
	// solution(3, 4, []int{1, 2, 3, 1, 2, 3, 1})
	solution(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2})
}
