package solution

func Solution(S string) (answer int) {
	if len(S) == 1 {
		return 0
	}

	if len(S)%2 == 0 {
		return -1
	}

	for i := 0; i < len(S)/2; i++ {
		if S[i] != S[len(S)-1-i] {
			return -1
		}
	}

	return len(S) / 2
}
