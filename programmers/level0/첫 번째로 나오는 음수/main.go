package main

func solution(num_list []int) (result int) {
	for i, v := range num_list {
		if v < 0 {
			return i
		}
	}
	if result == 0 {
		return -1
	}
	return
}
