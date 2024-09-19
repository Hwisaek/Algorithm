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

var (
	dp [][]int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ := strconv.Atoi(scan())
	dp = make([][]int, N+1)

	for i := 0; i < N; i++ {
		dp[i]
	}

	fmt.Fprint(w, f(int(N-1)))
}


1	1	1	2
2	1	2	3
3	1	2	4
4	1	2	4
5	1	2	4
6	1	2	4
7	1	2	4
8	1	2	3
9	1	1	2

