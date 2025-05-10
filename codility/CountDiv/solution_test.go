package solution

import (
	"testing"
)

// A <= i <= B 인 i 중에서 K로 나누어 떨어지는 수의 개수 찾기

// A, B 범위: 0 ~ 20억
// K 범위: 1 ~ 20억
// A <= B
func Solution(A int, B int, K int) (answer int) {
	answer += B/K + 1

	if A > 0 {
		answer -= (A-1)/K + 1
	}

	return
}

func TestSolution(t *testing.T) {
	type arg struct {
		A int
		B int
		K int
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
				6, 11, 2,
			},
			expect: expect{
				answer: 3,
			},
		},
		{
			name: "",
			arg: arg{
				0, 1, 2,
			},
			expect: expect{
				answer: 1,
			},
		},
		{
			name: "",
			arg: arg{
				0, 10, 2,
			},
			expect: expect{
				answer: 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.arg.A, tt.arg.B, tt.arg.K)
			if got != tt.expect.answer {
				t.Errorf("%s: Solution() = %v, want %v", tt.name, got, tt.expect.answer)
			}
		})
	}
}
