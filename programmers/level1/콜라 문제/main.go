package main

func solution(a int, b int, n int) (result int) {
	for n >= a {
		n = n - a + b
		result += b
	}
	return
}
