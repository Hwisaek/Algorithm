package main

import "math"

func solution(number int, limit int, power int) (count int) {
	for i := 1; i <= number; i++ {
		c := getCount(i)
		if c > limit {
			count += power
		} else {
			count += c
		}
	}
	return
}

func getCount(num int) (count int) {
	sqrt := int(math.Sqrt(float64(num)))
	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			// 약수를 찾으면 count를 2씩 증가
			count += 2
		}
	}

	if sqrt*sqrt == num {
		count--
	}
	return
}
