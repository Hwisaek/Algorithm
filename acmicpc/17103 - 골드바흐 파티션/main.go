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

	var T, m int
	Fscanln(r, &T)

	a := make([]int, T)
	for i := 0; i < T; i++ {
		Fscanln(r, &a[i])
		m = max(m, a[i])
	}

	isNotPrime := make([]bool, m+1)
	isNotPrime[0], isNotPrime[1] = true, true
	for i := 2; i*i <= m; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * 2; j <= m; j += i {
			isNotPrime[j] = true
		}
	}

	for _, e := range a {
		s := 0
		for i := 2; i*2 <= e; i++ {
			if !isNotPrime[i] && !isNotPrime[e-i] {
				s++
			}
		}
		Println(s)
	}
}
