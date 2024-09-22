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

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ := strconv.Atoi(scan())
	a := make([]int, 6)
	for i := range a {
		n, _ := strconv.Atoi(scan())
		a[i] = n
	}

	T, _ := strconv.Atoi(scan())
	P, _ := strconv.Atoi(scan())

	t := 0
	for _, e := range a {
		t += (e + T - 1) / T
	}
	fmt.Fprintln(w, t)

	fmt.Fprintln(w, N/P, N%P)
}
