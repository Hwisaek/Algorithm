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
	W, _ := strconv.Atoi(scan())
	H, _ := strconv.Atoi(scan())
	L, _ := strconv.Atoi(scan())

	fmt.Fprintf(w, "%d", min(N, (W/L)*(H/L)))
}
