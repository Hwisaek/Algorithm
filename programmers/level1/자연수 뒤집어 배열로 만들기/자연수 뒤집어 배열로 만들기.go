package 자연수_뒤집어_배열로_만들기

import "strconv"

func solution(n int64) (result []int) {
	itoa := strconv.Itoa(int(n))
	for _, c := range itoa {
		result = append([]int{int(c - '0')}, result...)
	}
	return
}
