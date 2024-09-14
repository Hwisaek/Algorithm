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
	col       []bool
	leftDiag  []bool
	rightDiag []bool
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())
	col = make([]bool, N)
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
		rightDiagIdx := rowIdx - colIdx + N
		leftDiagIdx := rowIdx + colIdx
		if col[rowIdx] || rightDiag[rightDiagIdx] || leftDiag[leftDiagIdx] {
			continue
		}

		col[rowIdx] = true
		rightDiag[rightDiagIdx] = true
		leftDiag[leftDiagIdx] = true

		dive(rowIdx + 1)

		col[rowIdx] = false
		rightDiag[rightDiagIdx] = false
		leftDiag[leftDiagIdx] = false
	}
}
