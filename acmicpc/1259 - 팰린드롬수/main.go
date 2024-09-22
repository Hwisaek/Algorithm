package main

import (
	"bufio"
	"fmt"
	"os"
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

l1:
	for {
		N := scan()
		if N == "0" {
			return
		}

		for i := 0; i < len(N)/2; i++ {
			if N[i] != N[len(N)-1-i] {
				fmt.Fprintln(w, "no")
				continue l1
			}
		}
		fmt.Fprintln(w, "yes")
	}
}
