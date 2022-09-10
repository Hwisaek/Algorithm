package _0

import "testing"

func isValid(s string) bool {

	m := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if a, ok := m[c]; ok {
			if len(stack) == 0 {
				return false
			}
			if a != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{s: "()"}, want: true},
		{name: "", args: args{s: "()[]{}"}, want: true},
		{name: "", args: args{s: "(]"}, want: false},
		{name: "", args: args{s: "["}, want: false},
		{name: "", args: args{s: "(("}, want: false},
		{name: "", args: args{s: "(){}}{"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
