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
	board = make([][]int, 9)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	for i := 0; i < 9; i++ {
		board[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			N, _ := strconv.Atoi(scan())
			board[i][j] = N
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
	row, col, found := findEmptyCell()
	if !found {
		print()
	}

	for i := 1; i <= 9; i++ {
		if check(row, col, i) {
			board[row][col] = i
			dive()
			board[row][col] = 0
		}
	}
}

func findEmptyCell() (row, col int, found bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}

	return
}

func check(row, col, num int) (flag bool) {
	var r, c, b int16

	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}
		n := board[row][i]
		r = r | 1<<n
	}
	if r&(1<<num) > 0 {
		return
	}

	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		n := board[i][col]
		c = c | 1<<n
	}
	if c&(1<<num) > 0 {
		return
	}

	startI := row / 3 * 3
	startJ := col / 3 * 3
	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			if i == col && j == row {
				continue
			}
			n := board[i][j]
			b = b | 1<<n
		}
	}
	if b&(1<<num) > 0 {
		return
	}

	return true
}
