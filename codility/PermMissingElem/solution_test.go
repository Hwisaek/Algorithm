package solution

func Solution(A []int) (answer int) {
	N := len(A)
	m := make(map[int]struct{})
	for i := 1; i <= N+1; i++ {
		m[i] = struct{}{}
	}

	for _, v := range A {
		delete(m, v)
	}

	for k := range m {
		return k
	}

	return
}
