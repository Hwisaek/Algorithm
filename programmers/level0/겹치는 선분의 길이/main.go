package main

import "math"

func solution(lines [][]int) (result int) {
	first, last := 101, -101
	for _, line := range lines {
		first = int(math.Min(float64(first), float64(line[0])))
		last = int(math.Max(float64(last), float64(line[1])))
	}

	count := make([]int, last-first)
	for _, line := range lines {
		start, end := line[0], line[1]
		for j := start - first; j < end-first; j++ {
			count[j]++
			if count[j] == 2 {
				result++
			}
		}
	}
	return
}

func main() {
	solution([][]int{{0, 1}, {2, 5}, {3, 9}})
	solution([][]int{{-1, 1}, {1, 3}, {3, 9}})
	solution([][]int{{0, 5}, {3, 9}, {1, 10}})
}
