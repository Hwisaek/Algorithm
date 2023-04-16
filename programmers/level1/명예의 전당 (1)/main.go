package main

import "sort"

func solution(k int, score []int) (result []int) {
	list := make([]int, 0)

	for i := 0; i < len(score); i++ {
		list = append(list, score[i])
		sort.Ints(list)
		if len(list) > k {
			list = list[1 : 1+k]
		}
		result = append(result, list[0])
	}

	return
}

func main() {
	solution(3, []int{10, 100, 20, 150, 1, 100, 200})
	solution(4, []int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000})
}
