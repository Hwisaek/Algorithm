func solution(n int) (result []int) {
	for i := 2; i <= n; {
		if n%i == 0 {
			n /= i
			if len(result) > 0 && result[len(result)-1] == i {
				continue
			}
			result = append(result, i)
		} else {
			i++
		}
	}

	return result
}
