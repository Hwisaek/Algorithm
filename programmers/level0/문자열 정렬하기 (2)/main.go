package main

import (
	"sort"
	"strings"
)

func solution(my_string string) string {
	my_string = strings.ToLower(my_string)
	runes := []rune(my_string)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
