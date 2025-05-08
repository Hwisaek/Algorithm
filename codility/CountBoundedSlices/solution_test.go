package solution

// K: max - min 의 최댓값
// K 에 해당하는 연속되는 A 의 부분 배열의 수

// 조건
// K: 0 ~ 10억
// N(A의 길이): 1 ~ 10만
// A의 원소의 범위: -10억 ~ 10억
func Solution(K int, A []int) (answer int) {
	N := len(A)

	for i := 0; i < N; i++ {
		MAX, MIN := A[i], A[i]

		for j := i; j < N; j++ {
			MAX = max(MAX, A[j])
			MIN = min(MIN, A[j])
			if MAX-MIN > K {
				break
			}
			answer++

			if answer > 1000000000 {
				return 1000000000
			}
		}
	}
	return
}
