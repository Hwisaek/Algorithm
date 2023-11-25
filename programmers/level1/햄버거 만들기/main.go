package main

import "fmt"

func solution(ingredient []int) (count int) {
	for i := 0; i < len(ingredient)-3; i++ {
		if ingredient[i] == 1 && ingredient[i+1] == 2 && ingredient[i+2] == 3 && ingredient[i+3] == 1 {
			ingredient = append(ingredient[:i], ingredient[i+4:]...)
			count++
			i -= 3
			if i < -1 {
				i = -1
			}
		}
	}

	return
}

func main() {
	fmt.Println(solution([]int{2, 1, 1, 2, 3, 1, 2, 3, 1}))
	fmt.Println(solution([]int{1, 3, 2, 1, 2, 1, 3, 1, 2}))
	fmt.Println(solution([]int{1, 2, 3, 2, 1}))
	fmt.Println(solution([]int{1, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1}))
}
