package main

func solution(array []int, n int) (count int) {
	for _, e := range array {
		if e == n {
			count++
		}
	}
	return
}
