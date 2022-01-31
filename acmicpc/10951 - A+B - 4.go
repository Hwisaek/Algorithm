package main

import "fmt"

func main() {
	var a, b int
	for cnt, _ := fmt.Scanln(&a, &b); cnt == 2; cnt, _ = fmt.Scanln(&a, &b) {
		fmt.Println(a + b)
	}
}
