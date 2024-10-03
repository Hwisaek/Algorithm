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

	p, _ := strconv.Atoi(scan())
	sum := 0

	x := make([]int, 0, p)
	for i := 0; i < p; i++ {
		a, _ := strconv.Atoi(scan())
		x = append(x, a)
		sum += a
	}

	a, _ := strconv.Atoi(scan())
	b, _ := strconv.Atoi(scan())

	r := 0
	startIdx := a % p
	if startIdx < 0 {
		startIdx += p
	}

	endIdx := b % p
	if endIdx < 0 {
		endIdx += p
	}
	if endIdx < startIdx {
		endIdx += p
	}

	for i := startIdx; i < endIdx; i++ {
		idx := i % p
		if idx < 0 {
			idx += p
		}

		r += x[idx]
	}

	r += sum * ((b - a) / p)

	fmt.Fprint(w, r)
}
