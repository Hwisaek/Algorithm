package solution

func Solution(X int, Y int, D int) (answer int) {
	target := Y - X
	answer = (target + D - 1) / D
	return
}
