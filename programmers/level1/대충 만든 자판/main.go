package main

func solution(keymap []string, targets []string) (result []int) {
	m := map[rune]int{}

	for _, s := range keymap {
		for i, c := range s {
			if _, ok := m[c]; ok {
				if m[c] > i+1 {
					m[c] = i + 1
				}
			} else {
				m[c] = i + 1
			}
		}
	}

	for _, target := range targets {
		count := 0
		for _, c := range target {
			if i, ok := m[c]; ok {
				count += i
			} else {
				count = -1
				break
			}
		}
		if count == 0 {
			count = -1
		}
		result = append(result, count)
	}
	return
}
