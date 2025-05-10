package solution

func Solution(A []int) (answer int) {
	m := make(map[int]int)
	for _, v := range A {
		m[v]++
	}

	for k, v := range m {
		if v%2 == 1 {
			answer = k
			break
		}
	}

	return
}
