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
	N      int           // 참가자 수
	minVal = math.MaxInt // 격차의 최솟값
	A      [][]int       // 시너지 배열
	member []bool        // 각 멤버의 소속팀
	aPower int           // a 팀의 power
	count  int
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

	member = make([]bool, N)
	member[0] = true
	dive(1)

	fmt.Fprintf(w, "%d", minVal)
	fmt.Fprintf(w, "\n%d", count)
}

func dive(memberCount int) {
	if minVal == 0 {
		return
	}

	if memberCount == N/2 {
		bPower := teamBPower()
		abs := aPower - bPower
		abs = max(abs, -abs)
		minVal = min(minVal, abs)

		count++
		return
	}

	for i := memberCount; i < N; i++ {
		if member[i] {
			continue
		}

		newAPower := 0
		for j := 0; j < N; j++ {
			if !member[j] {
				continue
			}
			newAPower += A[i][j]
			newAPower += A[j][i]
		}

		aPower += newAPower
		member[i] = true

		dive(memberCount + 1)

		aPower -= newAPower
		member[i] = false
	}
}

func teamBPower() (power int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j && !member[i] && !member[j] {
				power += A[i][j]
			}
		}
	}
	return
}
