package main

import (
	"strings"
)

func main() {
	solution("011")
}

func solution(numbers string) (answer int) {
	numSlice := strings.Split(numbers, "")

	for i := 1; i <= len(numbers); i++ {
		n := getPermutation(numSlice, i)
		_ = n
	}

	return
}
func getPermutation(sl []string, n int) (permutations [][]string) {
	if n <= 0 || n > len(sl) {
		return nil
	}

	// 재귀를 사용하여 순열을 구하는 함수
	var generate func(current []string)
	generate = func(current []string) {
		if len(current) == n {
			permutation := make([]string, len(current))
			copy(permutation, current)
			permutations = append(permutations, permutation)
			return
		}

		for i := 0; i < len(sl); i++ {
			// 이미 선택한 원소는 건너뛰기
			if contains(current, sl[i]) {
				continue
			}
			current = append(current, sl[i])
			generate(current)
			current = current[:len(current)-1]
		}
	}

	generate([]string{})
	return permutations
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
