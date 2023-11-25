package main

import "fmt"

func solution(my_string string, m int, c int) (result string) {
	for i := 0; i < len(my_string)/m; i++ {
		result += string(my_string[i*m+c-1])
	}

	return result
}

func main() {
	fmt.Println(solution("ihrhbakrfpndopljhygc", 4, 2))
}
