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
	s.Split(bufio.ScanWords)

	A, _ := strconv.Atoi(scan())
	B, _ := strconv.Atoi(scan())

	if A > 2*B || A < B+1 {
		fmt.Fprint(w, "NO")
		return
	}

	cnt := 0
	for ; A > B+1; A, B = A-2, B-1 {
		cnt++
	}

	fmt.Fprintf(w, "YES\n%d\n", cnt+1)
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(w, "aba")
	}

	fmt.Fprintf(w, "%sa", strings.Repeat("ab", B))
}
