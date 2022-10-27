package 가장_큰_수_찾기

func solution(array []int) []int {
	max, idx := -1, -1
	for i, n := range array {
		if n >= max {
			max, idx = n, i
		}
	}
	return []int{max, idx}
}
