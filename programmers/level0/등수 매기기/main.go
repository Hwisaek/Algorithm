package main

func solution(scores [][]int) []int {
	var avgScores []float64
	var ranks []int

	// 각 학생의 평균 점수를 계산하여 avgScores에 저장
	for _, score := range scores {
		avgScores = append(avgScores, float64(score[0]+score[1])/2.0)
	}

	// 평균 점수를 기준으로 등수 계산
	for i, score := range avgScores {
		rank := 1

		for j, otherScore := range avgScores {
			if i != j && score < otherScore {
				rank++
			}
		}

		ranks = append(ranks, rank)
	}

	return ranks
}
