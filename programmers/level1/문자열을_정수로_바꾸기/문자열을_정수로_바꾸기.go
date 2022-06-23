package 문자열을_정수로_바꾸기

import "strconv"

func solution(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
