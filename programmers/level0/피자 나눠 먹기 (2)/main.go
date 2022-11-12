package main

import "fmt"

func solution(n int) int {
	return lcm(n, 6) / 6
}

func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	fmt.Println(solution(6))
}
