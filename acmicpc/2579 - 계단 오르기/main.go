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
	stair []int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	t, _ := strconv.Atoi(scan())

	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(scan())
		stair = append(stair, n)
	}

	n := len(stair)
	dp := make([]int, n)

	if n >= 1 {
		dp[0] = stair[0]
	}
	if n >= 2 {
		dp[1] = stair[0] + stair[1]
	}
	if n >= 3 {
		dp[2] = max(stair[0]+stair[2], stair[1]+stair[2])
	}

	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2]+stair[i], dp[i-3]+stair[i-1]+stair[i])
	}

	fmt.Fprint(w, dp[n-1])
}
