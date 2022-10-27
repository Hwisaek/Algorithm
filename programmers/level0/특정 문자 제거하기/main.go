package 특정_문자_제거하기

import (
	"strings"
)

func solution(my_string string, letter string) string {
	return strings.ReplaceAll(my_string, letter, "")
}
