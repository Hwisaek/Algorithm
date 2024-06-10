package main

func main() {
	solution([]int{7, 9, 1, 1, 4})
}

func solution(elements []int) int {
	answer := map[int]struct{}{}

	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements); j++ {
			sum := 0
			for k := 0; k <= i; k++ {
				idx := j + k
				if idx >= len(elements) {
					idx -= len(elements)
				}
				sum += elements[idx]
			}
			answer[sum] = struct{}{}
		}
	}

	return len(answer)
}
