package _3__Roman_to_Integer

import "testing"

func romanToInt(s string) (result int) {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	last := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		c := string(s[i])
		n := m[c]
		if last > n && i != l-1 {
			n = n * -1
		}

		result += n
		last = m[c]
	}
	return
}

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
