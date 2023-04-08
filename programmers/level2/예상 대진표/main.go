package main

func solution(n int, a int, b int) int {
	round := 0
	for a != b {
		round++
		a = (a + 1) / 2
		b = (b + 1) / 2
	}
	return round
}
