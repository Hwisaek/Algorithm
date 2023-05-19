package main

import (
	. "fmt"
	"strings"
)

func main() {
	var N int
	Scan(&N)
	Print(strings.Repeat("long ", N/4) + "int")
}
