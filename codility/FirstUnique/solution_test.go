package solution

func Solution(A []int) (answer int) {
	m := map[int]int{}

	for i, e := range A {
		if _, ok := m[e]; !ok {
			m[e] = i
		} else {
			m[e] = -1
		}
	}

	idx := len(A)
	for k, v := range m {
		if v == -1 {
			continue
		}

		if v < idx {
			idx = v
			answer = k
		}
	}

	if idx == len(A) {
		answer = -1
	}

	return
}
