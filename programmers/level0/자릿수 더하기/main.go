package main

import "strconv"

func solution(n int) (result int) {
	for _, c := range strconv.Itoa(n) {
		result += int(c) - '0'
	}
	return
}
