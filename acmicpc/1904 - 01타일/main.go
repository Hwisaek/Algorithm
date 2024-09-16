package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func scan() string {
	s.Scan()
	return s.Text()
}

var (
	F []int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ := strconv.Atoi(scan())
	F = make([]int, max(3, N+1))

	fmt.Fprint(w, f(N))
}

func f(n int) int {
	if n <= 2 {
		return n
	}

	if F[n] != 0 {
		return F[n]
	}

	F[n] = (f(n-1) + f(n-2)) % 15746
	return F[n]
}
