package main

import (
	. "fmt"
)

func main() {
	for {
		var a, b int
		Scan(&a, &b)
		switch {
		case a == b:
			return
		case a%b == 0:
			Println("multiple")
		case b%a == 0:
			Println("factor")
		default:
			Println("neither")
		}
	}
}
