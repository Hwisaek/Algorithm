package main

func solution(brown int, yellow int) []int {
	w, h := brown/3, 2

	for {
		h++

		if (brown+yellow)%h != 0 {
			continue
		}

		w = (brown + yellow) / h

		if (h-2)*(w-2) != yellow {
			continue
		}

		break
	}
	return []int{w, h}
}

func main() {
	solution(10, 2)
	solution(8, 1)
	solution(24, 24)
}
