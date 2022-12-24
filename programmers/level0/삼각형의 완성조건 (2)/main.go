package main

func solution(sides []int) (count int) {
	min := 0
	if sides[0] > sides[1] {
		min = sides[0] - sides[1]
	} else {
		min = sides[1] - sides[0]
	}
	sum := sides[0] + sides[1]
	for i := min + 1; i < sum; i++ {
		count++
	}
	return
}

func main() {
	solution([]int{1, 2})
	solution([]int{3, 6})
	solution([]int{11, 7})
}
