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

	var n, m int
	Fscanln(r, &n)

	for i := 0; i < n; i++ {
		Fscanln(r, &m)
		for j := m; true; j++ {
			if isPrime(j) {
				Fprintln(w, j)
				break
			}
		}
	}
}

func isPrime(n int) bool {
	switch n {
	case 0, 1:
		return false
	default:
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
