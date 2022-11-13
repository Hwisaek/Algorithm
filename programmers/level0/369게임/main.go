package main

func solution(order int) (result int) {
	for ; order > 0; order /= 10 {
		switch order % 10 {
		case 3, 6, 9:
			result++
		}
	}
	return
}
