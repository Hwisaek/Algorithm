package main

func solution(n int, m int, section []int) (count int) {
	wall := make([]bool, n)
	for _, v := range section {
		wall[v-1] = true
	}

	for i, v := range wall {
		if v {
			for j := i; j < len(wall) && j < i+m; j++ {
				wall[j] = false
			}
			count++
		}
	}
	return
}

func main() {
	// solution(8, 4, []int{2, 3, 6})
	// solution(5, 4, []int{1, 3})
	solution(4, 1, []int{1, 2, 3, 4})
}
