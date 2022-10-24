package main

func solution(n int) (sum int) {
	for i := 2; i <= n; i += 2 {
		sum += i
	}
	return
}
