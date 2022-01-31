package main

import "fmt"

func main() {
	var score int
	grade := "F"

	_, _ = fmt.Scanln(&score)
	if score > 89 {
		grade = "A"
	} else if score > 79 {
		grade = "B"
	} else if score > 69 {
		grade = "C"
	} else if score > 59 {
		grade = "D"
	}
	fmt.Println(grade)
}
