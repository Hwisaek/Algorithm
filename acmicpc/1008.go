package main

import "fmt"

func main() {
	var A, B float64
	_, _ = fmt.Scanln(&A, &B)
	fmt.Println(A / B)
}
