package 최대공약수와_최소공배수

func solution(n int, m int) []int {
	return []int{gcd(n, m), lcm(n, m)}
}

func gcd(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return gcd(b, r)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
