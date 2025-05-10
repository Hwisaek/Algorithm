package solution

func Solution(A []int) (answer int) {
	lSum, rSum := A[0], 0
	for i := 1; i < len(A); i++ {
		rSum += A[i]
	}

	diff := abs(rSum - lSum)
	for i := 1; i < len(A)-1; i++ {
		lSum += A[i]
		rSum -= A[i]
		diff = min(diff, abs(rSum-lSum))
	}

	return diff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
