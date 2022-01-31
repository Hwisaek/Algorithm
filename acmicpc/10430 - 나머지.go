package main

import "fmt"

func main() {
	var A, B, C int
	_, _ = fmt.Scanln(&A, &B, &C)
	fmt.Println((A + B) % C)
	fmt.Println(((A % C) + (B % C)) % C)
	fmt.Println((A * B) % C)
	fmt.Println(((A % C) * (B % C)) % C)
}
