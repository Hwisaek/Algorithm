package main

import "fmt"

func main() {
	var a, b int
	for _, _ = fmt.Scanln(&a, &b); a != 0 && b != 0; _, _ = fmt.Scanln(&a, &b) {
		fmt.Println(a + b)
	}
}
