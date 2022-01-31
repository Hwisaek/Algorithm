package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanln(&a, &b)
	fmt.Println(a + b)
}
