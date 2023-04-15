package main

func solution(n int, times []int) int {
	left, right := 1, n*max(times)
	var answer int

	for left <= right {
		mid := (left + right) / 2
		cnt := 0

		for _, t := range times {
			cnt += mid / t
			if cnt >= n {
				break
			}
		}

		if cnt >= n {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return answer
}

func max(arr []int) int {
	var result int
	for i, a := range arr {
		if i == 0 || a > result {
			result = a
		}
	}
	return result
}
