package main

func solution(array []int, height int) (result int) {
	for _, e := range array {
		if e > height {
			result++
		}
	}
	return
}
