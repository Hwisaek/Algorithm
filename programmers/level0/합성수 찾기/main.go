package main

func solution(n int) (result int) {
	for i := 4; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
			if count >= 3 {
				result++
				break
			}
		}
	}
	return
}
