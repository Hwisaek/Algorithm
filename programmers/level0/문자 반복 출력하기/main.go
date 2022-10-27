package 문자_반복_출력하기

import "strings"

func solution(my_string string, n int) string {
	var b []rune
	for _, c := range my_string {
		b = append(b, []rune(strings.Repeat(string(c), n))...)
	}
	return string(b)
}
