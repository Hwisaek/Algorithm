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

	k, _ := strconv.Atoi(scan())
	a, _ := strconv.Atoi(scan())
	b, _ := strconv.Atoi(scan())
	c, _ := strconv.Atoi(scan())
	d, _ := strconv.Atoi(scan())

	if a*k+b == c*k+d {
		fmt.Fprint(w, "Yes ", a*k+b)
		return
	}
	fmt.Fprint(w, "No")
}
