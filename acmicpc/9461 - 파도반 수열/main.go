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
	F = make([]int, 101)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	T, _ := strconv.Atoi(scan())

	for i := 0; i < T; i++ {
		N, _ := strconv.Atoi(scan())

		fmt.Fprintln(w, f(N))
	}
}

func f(n int) int {
	switch n {
	case 1, 2, 3:
		return 1
	}

	if F[n] != 0 {
		return F[n]
	}
	F[n] = f(n-2) + f(n-3)
	return F[n]
}
