package 정수_제곱근_판별

import "math"

func solution(n int64) int64 {
	sqrt := math.Sqrt(float64(n))
	if sqrt == float64(int(sqrt)) {
		return int64(math.Pow(sqrt+1, 2))
	}
	return -1
}
