package solution

func Solution(A []int, X int) (answer int) {
	m := make(map[int]int)
	for _, e := range A {
		m[e]++
	}

	l := make([]int, 0)
	for k, v := range m {
		if v >= 2 {
			l = append(l, k)
		}
	}

	for i := 0; i < len(l); i++ {
		for j := i; j < len(l); j++ {
			a, b := l[i], l[j]

			if a == b {
				if m[a] >= 4 && a*b >= X {
					answer++
				}
			} else {
				if a*b >= X {
					answer++
				}
			}

			if answer > 1_000_000_000 {
				return -1
			}
		}
	}

	return
}
