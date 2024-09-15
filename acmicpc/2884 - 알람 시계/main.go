package main

import (
	"fmt"
	"time"
)

func main() {
	var h, m int

	_, _ = fmt.Scan(&h, &m)
	t := time.Date(0, 0, 0, h, m, 0, 0, time.Local).Add(-45 * time.Minute)
	fmt.Println(t.Hour(), t.Minute())
}
