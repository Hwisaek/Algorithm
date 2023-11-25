package main

func solution(num_list []int) int {
	var sum1, sum2 int

	for i, v := range num_list {
		if i%2 == 0 {
			sum1 += v
		} else {
			sum2 += v
		}
	}

	if sum1 > sum2 {
		return sum1
	}
	return sum2
}
