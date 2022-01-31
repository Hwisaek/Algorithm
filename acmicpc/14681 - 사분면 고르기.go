package main

import (
	"fmt"
)

func main() {
	var x, y int

	_, _ = fmt.Scan(&x, &y)
	switch {
	case x > 0 && y > 0:
		fmt.Println(1)
	case x < 0 && y > 0:
		fmt.Println(2)
	case x > 0 && y < 0:
		fmt.Println(4)
	case x < 0 && y < 0:
		fmt.Println(3)
	}
}
