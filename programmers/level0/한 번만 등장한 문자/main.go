package main

import "sort"

func solution(s string) string {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	runes := []rune{}
	for k, v := range m {
		if v == 1 {
			runes = append(runes, k)
		}
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
