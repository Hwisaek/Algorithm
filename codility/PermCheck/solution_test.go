package solution

import "testing"

func Solution(A []int) int {
	length := len(A)

	m := make(map[int]struct{})
	for _, e := range A {
		m[e] = struct{}{}
		if e > length {
			return 0
		}
	}

	if length == len(m) {
		return 1
	}

	return 0
}

func TestSolution(t *testing.T) {
	type arg struct {
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
				[]int{4, 1, 3, 2},
			},
			expect: expect{
				answer: 1,
			},
		},
		{
			name: "",
			arg: arg{
				[]int{4, 1, 3},
			},
			expect: expect{
				answer: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.arg.A)
			if got != tt.expect.answer {
				t.Errorf("%s: Solution() = %v, want %v", tt.name, got, tt.expect.answer)
			}
		})
	}
}
