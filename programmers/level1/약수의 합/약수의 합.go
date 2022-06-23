package

func solution(n int) (sum int) {
	if n == 0 || n == 1 {
		return n
	}

	divisors := map[int]struct{}{}

	for i := 1; i < n/2+1; i++ {
		if n%i == 0 {
			divisors[i] = struct{}{}
			divisors[n/i] = struct{}{}
		}
	}

	for divisor := range divisors {
		sum += divisor
	}

	return
}
