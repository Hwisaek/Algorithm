package main

func solution(keyinput []string, board []int) (loc []int) {
	loc = make([]int, 2)
	for _, key := range keyinput {
		switch {
		case key == "up" && loc[1] < board[1]/2:
			loc[1]++
		case key == "down" && loc[1] > -board[1]/2:
			loc[1]--
		case key == "left" && loc[0] > -board[0]/2:
			loc[0]--
		case key == "right" && loc[0] < board[0]/2:
			loc[0]++
		}
	}
	return
}

func main() {
	solution([]string{"left", "right", "up", "right", "right"}, []int{11, 11})
}
