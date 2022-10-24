package main

func solution(numList []int) (result []int) {
	result = make([]int, 2)
	for _, n := range numList {
		if n%2 == 0 {
			result[0]++
		} else {
			result[1]++
		}
	}
	return
}
