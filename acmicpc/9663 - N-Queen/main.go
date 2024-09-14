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
	N         = 0
	answer    = 0
	row       []bool
	leftDiag  []bool
	rightDiag []bool
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())
	row = make([]bool, N)
	leftDiag = make([]bool, N*2)
	rightDiag = make([]bool, N*2)

	for i := 0; i < N; i++ {
		dive(0)
	}

	fmt.Fprint(w, answer)
}

func dive(rowIdx int) {
	if rowIdx == N-1 {
		answer++
		return
	}

	for colIdx := 0; colIdx < N; colIdx++ {
		if row[rowIdx] {
			continue
		}

		rightDiagIdx := rowIdx - colIdx + N
		if rightDiag[rightDiagIdx] {
			continue
		}

		leftDiagIdx := rowIdx + colIdx
		if leftDiag[leftDiagIdx] {
			continue
		}

		row[rowIdx] = true
		rightDiag[rightDiagIdx] = true
		leftDiag[leftDiagIdx] = true

		dive(rowIdx + 1)

		row[rowIdx] = false
		rightDiag[rightDiagIdx] = false
		leftDiag[leftDiagIdx] = false
	}
}
