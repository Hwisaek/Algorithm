package main

func solution(rsp string) string {
	result := make([]rune, len(rsp))
	m := map[rune]rune{'2': '0', '0': '5', '5': '2'}
	for i, c := range rsp {
		result[i] = m[c]
	}
	return string(result)
}
