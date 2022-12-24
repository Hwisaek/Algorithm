package main

func solution(dots [][]int) int {
	x1, y1 := 256, 256
	x2, y2 := -256, -256
	for _, dot := range dots {
		x, y := dot[0], dot[1]
		if x < x1 {
			x1 = x
		}
		if x > x2 {
			x2 = x
		}
		if y < y1 {
			y1 = y
		}
		if y > y2 {
			y2 = y
		}
	}
	return (x2 - x1) * (y2 - y1)
}

func main() {
	solution([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}})
	solution([][]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}})
}
