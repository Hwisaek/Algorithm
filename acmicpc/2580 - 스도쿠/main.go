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
	board    = make([][]int, 9)
	row, col = make([]int, 9), make([]int, 9)
	box      = make([][]int, 3)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	for i := 0; i < 3; i++ {
		box[i] = make([]int, 3)
	}

	for i := 0; i < 9; i++ {
		board[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			N, _ := strconv.Atoi(scan())
			board[i][j] = N

			row[i] = row[i] | 1<<N
			col[j] = col[j] | 1<<N
			box[i/3][j/3] = box[i/3][j/3] | 1<<N
		}
	}

	dive()
}

func print() {
	for i := range board {
		for j := range board[i] {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}

	os.Exit(0)
}

func dive() {
	rowIdx, colIdx, found := findEmptyCell()
	if !found {
		print()
	}

	for i := 1; i <= 9; i++ {
		if check(rowIdx, colIdx, i) {
			board[rowIdx][colIdx] = i
			row[rowIdx] = row[rowIdx] | 1<<i
			col[colIdx] = col[colIdx] | 1<<i
			box[rowIdx/3][colIdx/3] = box[rowIdx/3][colIdx/3] | 1<<i

			dive()

			board[rowIdx][colIdx] = 0
			row[rowIdx] = row[rowIdx] ^ 1<<i
			col[colIdx] = col[colIdx] ^ 1<<i
			box[rowIdx/3][colIdx/3] = box[rowIdx/3][colIdx/3] ^ 1<<i
		}
	}
}

func findEmptyCell() (rowIdx, colIdx int, found bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}

	return
}

func check(rowIdx, colIdx, num int) (flag bool) {
	r := row[rowIdx]
	c := col[colIdx]
	b := box[rowIdx/3][colIdx/3]

	if r&(1<<num) > 0 {
		return
	}

	if c&(1<<num) > 0 {
		return
	}

	if b&(1<<num) > 0 {
		return
	}

	return true
}
