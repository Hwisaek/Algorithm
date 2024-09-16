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

	T, _ := strconv.Atoi(scan())

	for i := 0; i < T; i++ {
		N, _ := strconv.Atoi(scan())

		a, b, c := 1, 1, 1

		for i := 4; i <= N; i++ {
			a, b, c = b, c, a+b
		}

		fmt.Fprintln(w, c)
	}
}
