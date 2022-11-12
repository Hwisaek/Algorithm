package main

func solution(my_string string) (result int) {
	for _, c := range my_string {
		if '0' < c && c <= '9' {
			result += int(c - '0')
		}
	}
	return
}
