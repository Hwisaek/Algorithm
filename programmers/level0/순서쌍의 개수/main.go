package main

func solution(n int) (result int) {
	for i := 1; i <= n; i++ {
		if n/i*i == n {
			result++
		}
	}
	return
}
