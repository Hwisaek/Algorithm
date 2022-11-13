package main

func solution(my_string string, num1 int, num2 int) string {
	result := []rune(my_string)
	result[num1], result[num2] = result[num2], result[num1]
	return string(result)
}
