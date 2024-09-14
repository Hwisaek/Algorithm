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
			if board[i][j] != 0 {
				continue
			}
			N, _ := strconv.Atoi(scan())
			board[i][j] = N
		}
	}

	dive(0)

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

func dive(i int) {
	if i == 9 {
		print()
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				continue
			}

			for k := 1; k < 10; k++ {
				board[i][j] = k

				if !(checkRow(i, j) && checkCol(i, j) && checkBox(i, j)) {
					continue
				}

				dive(i + 1)
			}
		}
	}
}

func checkRow(i, j int) bool {
	m := map[int]bool{}
	for j := 0; j < 9; j++ {
		n := board[i][j]

		if m[n] {
			return false
		}

		m[n] = true
	}

	return true
}

func checkCol(i, j int) bool {
	m := map[int]bool{}
	for i := 0; i < 9; i++ {
		n := board[i][j]

		if m[n] {
			return false
		}

		m[n] = true
	}

	return true
}

func checkBox(i, j int) bool {
	m := map[int]bool{}

	startI := i / 3 * 3
	startJ := j / 3 * 3
	for i := startJ; i < startI+i; i++ {
		for j := startJ; j < startJ+3; j++ {
			n := board[i][j]

			if m[n] {
				return false
			}
			m[n] = true
		}
	}

	return true
}
