package solution

import (
	"fmt"
	"testing"
)

// K: 0 ~ 50
// C: 깨끗한 양말
// D: 더러운 양말
// C, D 의 각 원소: 1 ~ 50. 빈 배열이 아니고 최대 길이 50.
func Solution(K int, C []int, D []int) (answer int) {
	clean := map[int]int{}
	dirty := map[int]int{}

	// 1. C 에서 짝이 맞는 양말 제거
	for _, v := range C {
		clean[v]++
		if clean[v] == 2 {
			answer++
			delete(clean, v)
		}
	}

	for _, v := range D {
		dirty[v]++
	}

	for k := range clean {
		if K == 0 {
			break
		}

		if dirty[k] == 0 {
			continue
		}

		delete(clean, k)
		dirty[k]--
		if dirty[k] == 0 {
			delete(dirty, k)
		}

		K--
		answer++
	}

	if K == 0 {
		return
	}

	// 2. C 에서 짝이 안맞는 양말 색 확인
	// 3. D 에서 짝이 안맞는 양말을 찾기
	for k, v := range dirty {
		if K < 2 {
			break
		}

		if v >= 2 {
			a := min(v/2, K/2)
			answer += a
			K -= a * 2
			dirty[k] -= a * 2
		}
	}

	return
}

func TestSolution(t *testing.T) {
	fmt.Println(Solution(2, []int{1, 2, 1, 1}, []int{1, 4, 3, 2, 4}))
	fmt.Println(Solution(4, []int{1}, []int{1, 2, 2}))
	fmt.Println(Solution(3, []int{1}, []int{1, 2, 2}))
}

func TestMapDeleteDuringRange(t *testing.T) {
	m := map[int]int{1: 10, 2: 20, 3: 30, 4: 40}
	order := []int{}

	for k := range m {
		order = append(order, k)
		if k == 2 {
			delete(m, 2)
			delete(m, 3)
		}
	}

	t.Logf("순회 순서: %v", order)

	t.Logf("남은 map: %v", m)
}
