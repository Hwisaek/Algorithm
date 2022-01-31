package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanln(&A, &B)
	fmt.Println(A * B)
}
