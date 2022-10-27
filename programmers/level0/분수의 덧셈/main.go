package main

func solution(denum1 int, num1 int, denum2 int, num2 int) []int {
	d := denum1*num2 + denum2*num1
	n := num1 * num2
	g := gcd(d, n)

	return []int{d / g, n / g}
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
