package main

func solution(before string, after string) int {
	dict := map[rune]int{}
	for _, c := range before {
		dict[c]++
	}
	for _, c := range after {
		dict[c]--
		if dict[c] < 0 {
			return 0
		}
	}
	return 1
}
