package 바탕화면_정리

func solution(wallpaper []string) []int {
	minX, minY, maxX, maxY := -1, -1, -1, -1

	for i, str := range wallpaper {
		for j, c := range str {
			if c != '#' {
				continue
			}

			if minX == -1 {
				minX, minY = i, j
				maxX, maxY = i, j
			}

			if i < minX {
				minX = i
			}
			if j < minY {
				minY = j
			}
			if i > maxX {
				maxX = i
			}
			if j > maxY {
				maxY = j
			}
		}

	}
	return []int{minX, minY, maxX + 1, maxY + 1}
}
