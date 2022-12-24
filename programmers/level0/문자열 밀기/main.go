package main

func solution(A string, B string) (count int) {
	a := []byte(A)
	for i := 0; i < len(B); i++ {
		if string(a) == B {
			return
		}
		a = append(a[len(a)-1:], a[0:len(a)-1]...)
		count++
	}
	return -1
}

func main() {
	solution("hello", "ohell")
	solution("apple", "elppa")
}
