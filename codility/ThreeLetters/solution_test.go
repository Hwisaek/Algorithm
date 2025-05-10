package solution

import (
	"strings"
	"testing"
)

func Solution(A int, B int) string {
	a, b := max(A, B), min(A, B)
	C := min(a-b, 2)
	y := a - b - C
	x := b - y

	q, w, e := "ab", "aab", "a"
	if B > A {
		q, w, e = "ba", "bba", "b"
	}

	return strings.Repeat(q, x) + strings.Repeat(w, y) + strings.Repeat(e, C)
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "",
			input: []int{5, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Solution(test.input[0], test.input[1])
			t.Logf("Solution(%v) = %v", test.input, got)
		})
	}
}
