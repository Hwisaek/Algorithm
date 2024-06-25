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

	var N int
	Fscanln(r, &N)

	a := []int{}
	for i, x, y := 0, 0, 0; i < N; i++ {
		y = x
		Fscanln(r, &x)
		if i > 0 {
			a = append(a, x-y)
		}
	}

	b := a[0]
	for i := 1; i < len(a); i++ {
		b = gcd(b, a[i])
	}

	z := 0
	for i := 0; i < len(a); i++ {
		z += a[i]/b - 1
	}

	Fprintln(w, z)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
