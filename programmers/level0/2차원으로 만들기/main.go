package main

func solution(num_list []int, n int) (result [][]int) {
	for i := 0; i < len(num_list); i += n {
		result = append(result, []int{})
		for j := 0; j < n; j++ {
			result[i/n] = append(result[i/n], num_list[i+j])
		}
	}
	return
}
