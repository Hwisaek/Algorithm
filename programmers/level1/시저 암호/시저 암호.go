package 시저_암호

import "unicode"

func solution(s string, n int) (result string) {
	arr := make([]int32, 0, len(result))
	for _, code := range s {
		char := code
		if unicode.IsLetter(code) {
			char += int32(n)
			if unicode.IsUpper(code) {
				if char > 'Z' {
					char -= 26
				}
			} else {
				if char > 'z' {
					char -= 26
				}
			}
		}
		arr = append(arr, char)
	}
	return string(arr)
}
