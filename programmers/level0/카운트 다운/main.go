package main

func solution(start int, end_num int) (result []int) {
	for i := start; i >= end_num; i-- {
		result = append(result, i)
	}
	return
}
