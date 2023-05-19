package main

func solution(s string) (result []int) {
	m := make(map[rune]int)

	for i, v := range s {
		idx, ok := m[v]
		if !ok {
			result = append(result, -1)
		} else {
			result = append(result, i-idx)
		}
		m[v] = i
	}

	return
}

func main() {
	solution("banana")
}
