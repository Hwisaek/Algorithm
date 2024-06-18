package main

import (
	. "fmt"
)

func main() {
	var a, b, c, d, e, f int
	Scan(&a, &b, &c, &d, &e, &f)
	Print((b*f-c*e)/(b*d-a*e), (a*f-c*d)/(a*e-b*d))
}
