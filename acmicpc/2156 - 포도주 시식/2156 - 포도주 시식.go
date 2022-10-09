package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
)

func main() {
	var n int
	Fscan(r, &n)

	wine := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(r, &wine[i])
	}

	dp := make([]int, n)
	dp[0] = wine[0]
	if n > 1 {
		dp[1] = wine[0] + wine[1]
	}
	if n > 2 {
		dp[2] = max(dp[1], wine[2]+wine[0], wine[2]+wine[1])
	}
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+wine[i], dp[i-3]+wine[i]+wine[i-1])
	}

	Println(dp[n-1])
}

func max(ints ...int) (max int) {
	max = ints[0]
	for _, n := range ints[1:] {
		if n > max {
			max = n
		}
	}
	return
}
