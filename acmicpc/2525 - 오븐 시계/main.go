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
	C, _ := strconv.Atoi(scan())

	total := A*60 + B + C

	hour := total / 60 % 24
	min := total % 60

	fmt.Fprintf(w, "%d %d", hour, min)
}
