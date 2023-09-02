package main

func solution(arr [][]int) [][]int {
	w, h := len(arr[0]), len(arr)

	if w < h {
		for i := range arr {
			arr[i] = append(arr[i], make([]int, h-w)...)
		}
	} else if w > h {
		for i := 0; i < w-h; i++ {
			arr = append(arr, make([]int, w))
		}
	}
	return arr
}
