package solution

func Solution(S string) (answer int) {
	if len(S) == 0 {
		return -1
	}
	if len(S) == 1 {
		return 0
	}
	for i := 0; i < len(S)/2; i++ {
		if S[i] == S[len(S)-1-i] {
			answer++
		} else {
			break
		}
	}
	return
}
