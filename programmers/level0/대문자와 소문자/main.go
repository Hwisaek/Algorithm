package main

func solution(my_string string) string {
	b := []rune{}
	for _, c := range my_string {
		switch {
		case c <= 'Z':
			b = append(b, c-'A'+'a')
		default:
			b = append(b, c-'a'+'A')
		}
	}
	return string(b)
}
