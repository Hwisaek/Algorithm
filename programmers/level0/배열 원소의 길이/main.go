package main

func solution(strList []string) (result []int) {
	result = make([]int, len(strList))
	for i, s := range strList {
		result[i] = len(s)
	}
	return
}
