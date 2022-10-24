package main

func solution(angle int) int {
	switch {
	case angle == 180:
		return 4
	case angle > 90:
		return 3
	case angle == 90:
		return 2
	default:
		return 1
	}
}
