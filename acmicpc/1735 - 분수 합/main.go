package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c, d int
	Fscanln(r, &a, &b)
	Fscanln(r, &c, &d)

	x := a*d + b*c
	y := b * d
	g := gcd(x, y)
	Fprintln(w, x/g, y/g)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
