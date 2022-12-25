package main

func solution(num int, total int) (result []int) {
	first := (total - (num-1)*num/2) / num
	for i := 0; i < num; i++ {
		result = append(result, first+i)
	}
	return result
}
