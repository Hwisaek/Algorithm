package main

func solution(n int, numlist []int) []int {
	for i := len(numlist) - 1; i >= 0; i-- {
		if numlist[i]%n != 0 {
			numlist = append(numlist[:i], numlist[i+1:]...)
		}
	}
	return numlist
}

func main() {
	solution(12, []int{2, 100, 120, 600, 12, 12})
}
