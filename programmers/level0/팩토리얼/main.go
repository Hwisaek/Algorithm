package main

func solution(n int) (i int) {
	fact := 1
	for i = 2; fact <= n; i++ {
		fact *= i
	}
	return i - 2
}
