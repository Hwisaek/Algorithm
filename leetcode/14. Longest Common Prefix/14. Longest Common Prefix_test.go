package _4__Longest_Common_Prefix

import (
	"testing"
)

func longestCommonPrefix(strings []string) (result string) {
	m := map[string]int{}

	l := len(strings)

	for _, s := range strings {
		l2 := len(s)
		func() {

		}()
		for j := range s {
			substr := s[:l2-j]
			m[substr]++
			if m[substr] == l {
				return substr
			}
		}
	}
	return
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "",
			args: args{
				strings: []string{"flower", "flow", "flight"},
			},
			wantResult: "fl",
		},
		{
			name: "",
			args: args{
				strings: []string{"dog", "racecar", "car"},
			},
			wantResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := longestCommonPrefix(tt.args.strings); gotResult != tt.wantResult {
				t.Errorf("longestCommonPrefix() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
