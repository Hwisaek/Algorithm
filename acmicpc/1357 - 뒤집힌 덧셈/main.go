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

	X, _ := strconv.Atoi(scan())
	Y, _ := strconv.Atoi(scan())

	fmt.Fprint(w, rev(rev(X)+rev(Y)))
}

func rev(n int) int {
	s := strconv.Itoa(n)
	var b []byte
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	r, _ := strconv.Atoi(fmt.Sprintf("%s", b))
	return r
}
