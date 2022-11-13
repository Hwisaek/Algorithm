package main

func solution(array []int) (answer int) {
	for _, e := range array {
		for y := e; y > 0; y /= 10 {
			if y%10 == 7 {
				answer++
			}
		}
	}
	return
}
