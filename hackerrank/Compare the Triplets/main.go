package main

import (
	"fmt"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) (arr []int32) {
	arr = make([]int32, 2)
	for i := range a {
		if a[i] > b[i] {
			arr[0]++
		} else if a[i] < b[i] {
			arr[1]++
		}
	}
	return
}

func main() {
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
}
