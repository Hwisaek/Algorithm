package main

func solution(i int, j int, k int) (answer int) {
	for x := i; x <= j; x++ {
		for y := x; y > 0; y /= 10 {
			if y%10 == k {
				answer++
			}
		}
	}
	return
}

func main() {
	solution(1, 13, 1)
}
