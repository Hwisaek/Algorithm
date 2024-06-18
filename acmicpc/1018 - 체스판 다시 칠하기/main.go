package main

import (
	. "fmt"
)

func main() {
	var N, M int
	Scanln(&N, &M)
	w, b := make([][]int, N+1), make([][]int, N+1)
	// 누적합 배열 초기화
	for i := 0; i <= N; i++ {
		w[i] = make([]int, M+1)
		b[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		s := ""
		Scanln(&s)
		for j := 1; j <= M; j++ {
			w[i][j] = w[i][j-1] + w[i-1][j] - w[i-1][j-1]
			b[i][j] = b[i][j-1] + b[i-1][j] - b[i-1][j-1]

			if (i+j)%2 == 0 {
				if s[j-1] == 'W' {
					b[i][j]++
				} else {
					w[i][j]++
				}
			} else {
				if s[j-1] == 'B' {
					b[i][j]++
				} else {
					w[i][j]++
				}
			}
		}
	}

	min := 64
	for i := 1; i <= N-7; i++ {
		for j := 1; j <= M-7; j++ {
			s1 := w[i+7][j+7] - w[i-1][j+7] - w[i+7][j-1] + w[i-1][j-1]
			s2 := b[i+7][j+7] - b[i-1][j+7] - b[i+7][j-1] + b[i-1][j-1]

			if s1 < min {
				min = s1
			}
			if s2 < min {
				min = s2
			}
		}
	}

	Print(min)
}
