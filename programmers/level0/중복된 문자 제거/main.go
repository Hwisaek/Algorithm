package main

func solution(my_string string) string {
	m := map[rune]interface{}{}

	runes := []rune{}
	for _, c := range my_string {
		if _, ok := m[c]; ok {
			continue
		}
		m[c] = struct{}{}
		runes = append(runes, c)
	}
	return string(runes)
}
