package main

import (
	"fmt"
)

func main() {
	var h, m int

	_, _ = fmt.Scan(&h, &m)
	if m < 45 {
		h -= 1
		m = 60 + m - 45
	}
	if h < 0 {
		h = 23
	}
	fmt.Println(h, m)
}
