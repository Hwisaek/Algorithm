package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	for {
		if !s.Scan() {
			break
		}
		N, _ := strconv.Atoi(s.Text())
		N = int(math.Pow(3, float64(N)))
		a(N)
		fmt.Fprintln(w)
	}
	w.Flush()
}

func a(i int) {
	if i <= 1 {
		fmt.Fprint(w, "-")
		return
	}

	a(i / 3)
	fmt.Fprint(w, strings.Repeat(" ", i/3))
	a(i / 3)
}
