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
	M, _ := strconv.Atoi(scan())
	K, _ := strconv.Atoi(scan())

	if K >= N+M-1 {
		fmt.Fprintln(w, "YES")
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				fmt.Fprintf(w, "%d ", i+j+1)
			}
			fmt.Fprintln(w)
		}
		return
	}

	fmt.Fprint(w, "NO")
}
