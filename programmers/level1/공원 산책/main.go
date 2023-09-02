package main

func solution(park []string, routes []string) []int {
	var r, c int

	for i, s := range park {
		for j, char := range s {
			if char == 'S' {
				r, c = i, j
			}
		}
	}

	maxR, maxC := len(park)-1, len(park[0])-1

for1:
	for _, route := range routes {
		dir, d := route[0], int(route[2]-'0')

		nextR, nextC := r, c
		for i := 0; i < d; i++ {
			switch dir {
			case 'N':
				nextR = nextR - 1
			case 'S':
				nextR = nextR + 1
			case 'W':
				nextC = nextC - 1
			case 'E':
				nextC = nextC + 1
			}

			if nextC > maxC || nextC < 0 || nextR > maxR || nextR < 0 {
				continue for1
			}

			if park[nextR][nextC] == 'X' {
				continue for1
			}
		}
		r, c = nextR, nextC
	}
	return []int{r, c}
}

func main() {
	//solution([]string{"SOO", "OOO", "OOO"}, []string{"E 2", "S 2", "W 1"})
	solution([]string{
		"SOO",
		"OXX",
		"OOO"}, []string{"E 2", "S 2", "W 1"})
	solution([]string{"OSO", "OOO", "OXO", "OOO"}, []string{"E 2", "S 3", "W 1"})
	solution([]string{"OSO", "OOO", "OXO", "OOO"}, []string{"E 9", "S 9", "W 9"})
}
