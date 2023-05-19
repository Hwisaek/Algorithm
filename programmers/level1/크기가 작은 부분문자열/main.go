package main

func solution(t string, p string) (count int) {
	l := len(p)
	for i := 0; i <= len(t)-l; i++ {
		if t[i:i+l] <= p {
			count++
		}
	}
	return
}

func main() {
	solution("10203", "15")
}
