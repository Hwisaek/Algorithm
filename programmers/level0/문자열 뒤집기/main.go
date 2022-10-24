package main

func solution(my_string string) (result string) {
	for i := len(my_string) - 1; i >= 0; i-- {
		result += string(my_string[i])
	}
	return
}
