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

	var a byte

	N, _ := strconv.Atoi(scan())
	for i := 0; i < N; i++ {
		s := scan()
		if a == 0 {
			a = s[0]
			continue
		}

		if s[0] != a {
			fmt.Fprint(w, 0)
			return
		}
	}

	fmt.Fprint(w, 1)
}
