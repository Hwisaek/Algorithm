package 정수_내림차순으로_배치하기

import (
	"math"
	"sort"
	"strconv"
)

func solution(n int64) (result int64) {
	itoa := strconv.Itoa(int(n))
	intArr := make([]int, 0, len(itoa))
	for _, c := range itoa {
		num := c - '0'
		intArr = append(intArr, int(num))
	}
	
  sort.Sort(sort.IntSlice(intArr))

	sum := 0
	for i, i2 := range intArr {
		sum += i2 * int(math.Pow(float64(10), float64(i)))
	}
	return int64(sum)
}
