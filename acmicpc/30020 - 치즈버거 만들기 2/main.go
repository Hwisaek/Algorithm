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

	n, _ := strconv.Atoi(scan())

	a := 180 * (n - 1)
	for i := 0; i < n; i++ {
		t, _ := strconv.Atoi(scan())
		a -= t * 2
	}

	fmt.Fprint(w, a)
}
