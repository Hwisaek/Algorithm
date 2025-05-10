package solution

import "testing"

// A, B, C 로 이뤄진 문자열 S 에서 AA, BB, CC 를 모두 제거하고 남은 문자열을 반환하시오

// N: 문자열 S의 길이.  0 ~ 5만
func Solution(S string) string {
	s := make([]byte, 0, len(S))

	for i := 0; i < len(S); i++ {
		s = append(s, S[i])

		if len(s) < 2 {
			continue
		}

		if s[len(s)-1] == s[len(s)-2] {
			s = s[:len(s)-2]
		}
	}

	return string(s)
}

func TestSolution(t *testing.T) {
	type arg struct {
		S string
	}
	type expect struct {
		answer string
	}

	tests := []struct {
		name   string
		arg    arg
		expect expect
	}{
		{
			name: "",
			arg: arg{
				"ACCAABBC",
			},
			expect: expect{
				answer: "AC",
			},
		},
		{
			name: "",
			arg: arg{
				"ABCBBCBA",
			},
			expect: expect{
				answer: "",
			},
		},
		{
			name: "",
			arg: arg{
				"BABABA",
			},
			expect: expect{
				answer: "BABABA",
			},
		},
		{
			name: "",
			arg: arg{
				"",
			},
			expect: expect{
				answer: "",
			},
		},
		{
			name: "",
			arg: arg{
				"A",
			},
			expect: expect{
				answer: "A",
			},
		},
		{
			name: "",
			arg: arg{
				"ABCCCBA",
			},
			expect: expect{
				answer: "ABCBA",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.arg.S)
			if got != tt.expect.answer {
				t.Errorf("%s: Solution() = %v, want %v", tt.name, got, tt.expect.answer)
			}
		})
	}
}
