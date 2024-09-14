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
	col = make([]bool, N)         // 각 열에 퀸이 있는지 여부
	leftDiag = make([]bool, 2*N)  // 왼쪽 위 대각선 (인덱스: row + col)
	rightDiag = make([]bool, 2*N) // 오른쪽 위 대각선 (인덱스: row - col + N)

	for i := 0; i < N; i++ {
		dive(0)
	}

	fmt.Fprint(w, answer)
}

func dive(rowIdx int) {
	if rowIdx == N {
		// 모든 행에 퀸을 배치한 경우, 하나의 유효한 솔루션을 찾았음
		answer++
		return
	}

	for colIdx := 0; colIdx < N; colIdx++ {
		// 현재 열과 대각선이 사용 중인지 확인
		leftDiagIdx := rowIdx + colIdx
		rightDiagIdx := rowIdx - colIdx + N
		if col[rowIdx] || leftDiag[leftDiagIdx] || rightDiag[rightDiagIdx] {
			continue
		}

		// 현재 위치에 퀸을 배치
		col[colIdx] = true
		leftDiag[leftDiagIdx] = true
		rightDiag[rightDiagIdx] = true

		// 다음 행으로 재귀 호출
		dive(rowIdx + 1)

		// 상태 복원
		col[colIdx] = false
		leftDiag[leftDiagIdx] = false
		rightDiag[rightDiagIdx] = false
	}
}
