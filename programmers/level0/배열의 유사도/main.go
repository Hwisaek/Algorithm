package main

func solution(s1 []string, s2 []string) (result int) {
	m := map[string]interface{}{}
	for _, s := range s1 {
		m[s] = struct{}{}
	}
	for _, s := range s2 {
		if _, ok := m[s]; ok {
			result++
		}
	}
	return
}
