package 안전지대

func solution(board [][]int) (result int) {
	width, height := len(board[0]), len(board)
	result = height * width
	dx := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dy := []int{1, 1, 0, -1, -1, -1, 0, 1}

	for r, row := range board {
		for c, e := range row {
			if e == 1 {
				result--
				for i := 0; i < 8; i++ {
					x, y := r+dx[i], c+dy[i]
					if x < 0 || x >= width || y < 0 || y >= height {
						continue
					}

					if board[x][y] == 0 {
						board[x][y] = 2
						result--
					}
				}
			}
		}
	}

	return
}
