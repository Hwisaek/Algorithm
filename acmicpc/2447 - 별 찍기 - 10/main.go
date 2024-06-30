package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	N, _ := strconv.Atoi(scan())
	fmt.Fprint(w, strings.Join(a(N), "\n"))
}

func a(i int) []string {
	if i == 1 {
		return []string{"*"}
	}

	stars := a(i / 3)
	l := make([]string, 0)

	for _, e := range stars {
		l = append(l, strings.Repeat(e, 3))
	}
	for _, e := range stars {
		l = append(l, e+strings.Repeat(" ", i/3)+e)
	}
	for _, e := range stars {
		l = append(l, strings.Repeat(e, 3))
	}

	return l
}
