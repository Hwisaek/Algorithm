package solution

import (
	"testing"
)

// 각 원소의 인덱스를 쫓아가면서 -1 이 나타날때까지 카운트
func Solution(A []int) (answer int) {
	for i := 0; i != -1; i = A[i] {
		answer++

	}
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "예제 테스트 케이스",
			input:    []int{1, 4, -1, 3, 2},
			expected: 4,
		},
		{
			name:     "짧은 리스트",
			input:    []int{1, -1, -1, 3, 2},
			expected: 2,
		},
		{
			name:     "긴 리스트",
			input:    []int{1, 2, 3, 4, -1},
			expected: 5,
		},
		{
			name:     "다른 리스트 구조",
			input:    []int{-1, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "순환 없는 더 긴 리스트",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Solution(tc.input)
			if result != tc.expected {
				t.Errorf("테스트 실패 %s: 입력: %v, 결과: %d, 기대값: %d",
					tc.name, tc.input, result, tc.expected)
			}
		})
	}
}
