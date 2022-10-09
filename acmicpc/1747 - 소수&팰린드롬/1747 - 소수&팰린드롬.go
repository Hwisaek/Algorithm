package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	maxN := 1003002

	var n int
	fmt.Scanf("%d", &n)

	isNotPrime := make([]bool, maxN+1)
	isNotPrime[0], isNotPrime[1] = true, true
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(maxN)))); i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * 2; j <= maxN; j += i {
			isNotPrime[j] = true
		}
	}

	for i := n; i <= maxN; i++ {
		if i >= n && !isNotPrime[i] && isPalindrome(i) {
			fmt.Println(i)
			return
		}
	}
}

func isPalindrome(n int) (flag bool) {
	s := strconv.Itoa(n)
	if s == reverse(s) {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, c := range s {
		runes[len(s)-1-i] = c
	}
	return string(runes)
}
