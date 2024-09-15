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
	A      [][]int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())

	A = make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			a, _ := strconv.Atoi(scan())
			A[i][j] = a
		}
	}

	// 모든 조합을 비트마스크로 탐색합니다.
	for mask := 0; mask < (1 << N); mask++ {
		if popCount(mask) == N/2 {
			aPower := 0
			bPower := 0

			// 팀 A의 파워 계산
			for i := 0; i < N; i++ {
				for j := i + 1; j < N; j++ {
					if (mask&(1<<i) > 0) && (mask&(1<<j) > 0) {
						aPower += A[i][j] + A[j][i]
					}
				}
			}

			// 팀 B의 파워 계산
			for i := 0; i < N; i++ {
				for j := i + 1; j < N; j++ {
					if (mask&(1<<i) == 0) && (mask&(1<<j) == 0) {
						bPower += A[i][j] + A[j][i]
					}
				}
			}

			// 최소 차이 계산
			abs := aPower - bPower
			abs = max(abs, -abs)
			minVal = min(minVal, abs)
		}
	}

	fmt.Fprintf(w, "%d", minVal)
}

// popCount는 비트마스크에서 1의 개수를 세는 함수입니다.
func popCount(mask int) int {
	count := 0
	for mask > 0 {
		count += mask & 1
		mask >>= 1
	}
	return count
}
