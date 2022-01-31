package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanln(&A, &B)
	switch {
	case A < B:
		fmt.Println("<")
	case A > B:
		fmt.Println(">")
	case A == B:
		fmt.Println("==")
	}
}
