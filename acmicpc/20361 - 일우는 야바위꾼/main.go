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

	_ = scan()
	X := scan()
	K, _ := strconv.Atoi(scan())

	for i := 0; i < K; i++ {
		A := scan()
		B := scan()

		switch X {
		case A:
			X = B
		case B:
			X = A
		}
	}

	fmt.Fprint(w, X)
}
