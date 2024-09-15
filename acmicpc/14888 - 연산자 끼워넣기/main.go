package main

import (
	"bufio"
	"fmt"
	"math"
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
	N      int
	minVal = math.MaxInt
	maxVal = math.MinInt
	A      []int
	B      = make([]int, 4)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())
	A = make([]int, N)
	for j := 0; j < N; j++ {
		a, _ := strconv.Atoi(scan())
		A[j] = a
	}

	for i := 0; i < 4; i++ {
		count, _ := strconv.Atoi(scan())
		B[i] = count
	}
	dive(1, A[0])

	fmt.Fprintf(w, "%d\n%d", maxVal, minVal)
}

func dive(i int, total int) {
	if i == N {
		minVal = min(minVal, total)
		maxVal = max(maxVal, total)
		return
	}

	for j := 0; j < 4; j++ {
		if B[j] == 0 {
			continue
		}
		B[j]--

		switch j {
		case 0:
			dive(i+1, total+A[i])
		case 1:
			dive(i+1, total-A[i])
		case 2:
			dive(i+1, total*A[i])
		case 3:
			dive(i+1, total/A[i])
		}

		B[j]++
	}
}
