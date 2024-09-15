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
	A = make([][][]int, 21)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	for i := 0; i < 21; i++ {
		A[i] = make([][]int, 21)
		for j := 0; j < 21; j++ {
			A[i][j] = make([]int, 21)
			for k := 0; k < 21; k++ {
				A[i][j][k] = -1

			}
		}
	}

	for {
		a, _ := strconv.Atoi(scan())
		b, _ := strconv.Atoi(scan())
		c, _ := strconv.Atoi(scan())

		if a == b && b == c && c == -1 {
			return
		}

		fmt.Fprintf(w, "w(%d, %d, %d) = %d\n", a, b, c, W(a, b, c))
	}
}

func W(a, b, c int) int {

	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}

	if a > 20 || b > 20 || c > 20 {
		return 1048576
	}

	if A[a][b][c] != -1 {
		return A[a][b][c]
	}

	if a < b && b < c {
		A[a][b][c] = W(a, b, c-1) + W(a, b-1, c-1) - W(a, b-1, c)

	} else {
		A[a][b][c] = W(a-1, b, c) + W(a-1, b-1, c) + W(a-1, b, c-1) - W(a-1, b-1, c-1)

	}

	return A[a][b][c]
}
