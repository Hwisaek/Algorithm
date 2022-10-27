package 모음_제거

import "regexp"

func solution(my_string string) string {
	return regexp.MustCompile("[aeiou]").ReplaceAllString(my_string, "")
}
