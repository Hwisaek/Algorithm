package main

import (
	"fmt"
)

func main() {
	var year int

	_, _ = fmt.Scan(&year)
	result := 0
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		result = 1
	}
	fmt.Println(result)
}
