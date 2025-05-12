package solution

import (
	"testing"
)

func Solution(X int, A []int) (answer int) {
	leaf := make(map[int]struct{})

	for i, e := range A {
		leaf[e] = struct{}{}

		if len(leaf) == X {
			return i
		}
	}

	return -1
}

func TestSolution(t *testing.T) {
	type arg struct {
		X int
		A []int
	}
	type expect struct {
		answer int
	}

	tests := []struct {
		name   string
		arg    arg
		expect expect
	}{
		{
			name: "",
			arg: arg{
				5,
				[]int{1, 3, 1, 4, 2, 3, 5, 4},
			},
			expect: expect{
				answer: 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.arg.X, tt.arg.A)
			if got != tt.expect.answer {
				t.Errorf("%s: Solution() = %v, want %v", tt.name, got, tt.expect.answer)
			}
		})
	}
}
