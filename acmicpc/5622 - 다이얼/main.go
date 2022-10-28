package main

import "fmt"

var m = []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}

func main() {
	var word string
	_, _ = fmt.Scanln(&word)

	time := 0
	for _, i2 := range word {
		time += m[i2-'A']
	}
	fmt.Println(time)
}
