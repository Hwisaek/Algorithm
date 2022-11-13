package main

func solution(hp int) (result int) {
	result += hp / 5
	hp %= 5
	result += hp / 3
	result += hp % 3
	return
}

func main() {
	solution(999)
}
