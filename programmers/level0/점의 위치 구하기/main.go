package 점의_위치_구하기

func solution(dot []int) (result int) {
	x, y := dot[0], dot[1]
	switch {
	case x > 0 && y > 0:
		return 1
	case x < 0 && y > 0:
		return 2
	case x < 0 && y < 0:
		return 3
	case x > 0 && y < 0:
		return 4
	}
	return
}
