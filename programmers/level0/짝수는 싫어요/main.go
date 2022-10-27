package 짝수는_싫어요

func solution(n int) (arr []int) {
	for i := 1; i <= n; i += 2 {
		arr = append(arr, i)
	}
	return
}
