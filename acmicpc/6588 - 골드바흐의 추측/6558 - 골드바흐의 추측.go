package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
)

const max = 1000000

var (
	r = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	isNotPrime := [max + 1]bool{true, true}
	for i := 2; i <= int(math.Sqrt(max)); i++ {
		if !isNotPrime[i] {
			for j := i * 2; j <= max; j += i {
				isNotPrime[j] = true
			}
		}
	}

	for r.Scan() {
		n, _ := strconv.Atoi(r.Text())
		if n == 0 {
			break
		}

		for i := 2; i <= n/2; i++ {
			if !isNotPrime[i] {
				if !isNotPrime[n-i] {
					Fprintf(w, "%d = %d + %d\n", n, i, n-i)
					break
				}
			}
		}
	}
	w.Flush()
}
