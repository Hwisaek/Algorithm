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
	f    []int
	a, b int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	n, _ := strconv.Atoi(scan())
	f = make([]int, n+1)

	fib(n)
	fibonacci(n)

	fmt.Fprintf(w, "%d %d", a, b+1)
}

func fib(n int) int {
	if n == 1 || n == 2 {
		a++
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibonacci(n int) int {
	f[1] = 1
	f[2] = 1
	for i := 3; i < n; i++ {
		b++
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
