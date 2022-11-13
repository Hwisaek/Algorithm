package main

func solution(cipher string, code int) (result string) {
	for i := code - 1; i < len(cipher); i += code {
		result += string(cipher[i])
	}
	return
}
