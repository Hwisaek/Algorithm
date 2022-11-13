package main

import "strconv"

func solution(num int, k int) int {
	str := strconv.Itoa(num)
	for i, c := range str {
		if int(c-'0') == k {
			return i + 1
		}
	}
	return -1
}
