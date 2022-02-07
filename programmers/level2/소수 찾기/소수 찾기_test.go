package 소수_찾기

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var list = make(map[int]bool)

func solution(numbers string) (result int) {
	primes := map[int]struct{}{}
	var numSlice []int
	for _, number := range numbers {
		numSlice = append(numSlice, int(number-'0'))
	}
	for i := 1; i <= len(numSlice); i++ {
		perms := permutation(numSlice, i)
		for _, perm := range perms {
			num, _ := strconv.Atoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), ""), "[]"))
			primes[num] = struct{}{}
		}
	}

	for num := range primes {
		if isPrime(num) {
			result++
		}
	}
	return
}

func isPrime(num int) (result bool) {
	return
}

func permutation(arr []int, r int) (result [][]int) {
	sort.Ints(arr)
	used := make([]int, len(arr))
	generate([]int{}, used, r, &arr, &result)
	return result
}

func generate(chosen []int, used []int, r int, arr *[]int, result *[][]int) {
	if len(chosen) == r {
		copySlice := make([]int, len(chosen))
		copy(copySlice, chosen)
		*result = append(*result, copySlice)
		return
	}

	for i := 0; i < len(*arr); i++ {
		if used[i] == 0 {
			chosen = append(chosen, (*arr)[i])
			used[i] = 1
			generate(chosen, used, r, arr, result)
			used[i] = 0
			chosen = chosen[:len(chosen)-1]
		}
	}
}

func Test_solution(t *testing.T) {
	type args struct {
		numbers string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{numbers: "17"},
			want: 3,
		},
		{
			name: "2",
			args: args{numbers: "011"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.numbers); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
