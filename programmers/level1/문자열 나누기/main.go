package main

func solution(s string) (result int) {
	m := make(map[bool]int)

	last := rune(s[0])
	for i, r := range s {
		m[last == r]++
		if m[true] == m[false] || len(s)-1 == i {
			result++
			m = make(map[bool]int)
			if len(s) > i+1 {
				last = rune(s[i+1])
			}
		}

	}
	return
}

func main() {
	// solution("banana")
	solution("abracadabra")
}
