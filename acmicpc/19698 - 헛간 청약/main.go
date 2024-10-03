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

	A, _ := strconv.Atoi(scan())
	B, _ := strconv.Atoi(scan())
	m, _ := strconv.Atoi(scan())

	s := int64(0)

	for i := 0; i < m; i++ {
		n, _ := strconv.Atoi(scan())

		a := int64(1)
		for j := 0; j < m-1-i; j++ {
			a *= int64(A)
		}

		s += int64(n) * a
	}

	for _, i := range strconv.FormatInt(s, B) {
		n, _ := strconv.ParseInt(string(i), B, 64)
		fmt.Fprintf(w, "%d ", n)
	}
}
