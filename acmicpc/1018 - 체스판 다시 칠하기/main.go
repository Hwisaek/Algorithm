package main

import (
	"bufio"
	. "fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var N, M int
	Fscanln(r, &N, &M)
	b := make([][]int, N+1)
	// 누적합 배열 초기화
	for i := 0; i <= N; i++ {
		b[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		s := ""
		Fscanln(r, &s)
		for j := 1; j <= M; j++ {
			b[i][j] = b[i][j-1] + b[i-1][j] - b[i-1][j-1]

			if (i+j)%2 == 0 {
				if s[j-1] == 'W' {
					b[i][j]++
				}
			} else {
				if s[j-1] == 'B' {
					b[i][j]++
				}
			}
		}
	}

	min := 64
	for i := 1; i <= N-7; i++ {
		for j := 1; j <= M-7; j++ {
			s2 := b[i+7][j+7] - b[i-1][j+7] - b[i+7][j-1] + b[i-1][j-1]

			if s2 < min {
				min = s2
			}
			if 64-s2 < min {
				min = 64 - s2
			}
		}
	}

	Print(min)
}
