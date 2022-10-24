package main

import "strings"

func solution(str1 string, str2 string) int {
	if strings.Contains(str1, str2) {
		return 1
	}
	return 2
}
