package main

import "sort"

func solution(my_string string) (result []int) {
	result = []int{}
	for _, c := range my_string {
		if '0' <= c && c <= '9' {
			result = append(result, int(c-'0'))
		}
	}
	sort.Ints(result)
	return
}
