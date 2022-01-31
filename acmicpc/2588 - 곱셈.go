package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanln(&A)
	_, _ = fmt.Scanln(&B)

	fmt.Println(A * (B % 10))
	fmt.Println(A * (B % 100 / 10))
	fmt.Println(A * (B % 1000 / 100))
	fmt.Println(A * B)
}
