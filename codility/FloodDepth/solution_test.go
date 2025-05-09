package solution

import "testing"

// N: 1 ~ 10만
// A의 원소: 1 ~ 1억
func Solution(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = A[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], A[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], A[i])
	}

	maxDepth := 0
	for i := 0; i < n; i++ {
		waterHeight := min(leftMax[i], rightMax[i])
		depth := waterHeight - A[i]
		maxDepth = max(maxDepth, depth)
	}

	return maxDepth
}

func TestSolution(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "",
			input: []int{3, 2, 1},
			want:  0,
		},
		{
			name:  "",
			input: []int{1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2},
			want:  2,
		},
		{
			name:  "",
			input: []int{5, 8},
			want:  0,
		},
		{
			name:  "",
			input: []int{5, 1, 6, 2, 7},
			want:  4,
		},
		{
			name:  "",
			input: []int{9, 1, 5, 2, 8},
			want:  7,
		},
		{
			name:  "",
			input: []int{3, 0, 2, 0, 4},
			want:  3,
		},
		{
			name:  "",
			input: []int{3, 0, 1, 0, 2, 0, 4},
			want:  3,
		},
		{
			name:  "",
			input: []int{20, 4, 10, 1, 3},
			want:  6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Solution(c.input)
			if got != c.want {
				t.Errorf("Solution(%v) = %d; want %d", c.input, got, c.want)
			}
		})
	}
}
