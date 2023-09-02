package main

import "strings"

func solution(my_string string, is_prefix string) int {
	if strings.HasPrefix(my_string, is_prefix) {
		return 1
	}
	return 0
}
