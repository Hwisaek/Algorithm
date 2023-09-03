package main

func solution(my_string string, n int) string {
	return string([]byte(my_string)[len(my_string)-n:])
}
