package main

func solution(name []string, yearning []int, photo [][]string) (result []int) {
	score := make(map[string]int)
	for i := range name {
		score[name[i]] = yearning[i]
	}

	for _, list := range photo {
		point := 0
		for _, v := range list {
			point += score[v]
		}
		result = append(result, point)
	}
	return
}
