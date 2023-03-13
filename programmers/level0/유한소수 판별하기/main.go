package main

func solution(a int, b int) int {
	remainder := a % b
	remainders := make(map[int]bool)

	for remainder != 0 && !remainders[remainder] {
		remainders[remainder] = true
		remainder = (remainder * 10) % b
	}

	if remainder == 0 {
		return 1
	} else {
		return 2
	}
}
