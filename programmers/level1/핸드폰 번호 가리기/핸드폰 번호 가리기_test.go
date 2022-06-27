package 핸드폰_번호_가리기

import "testing"

func solution(phone_number string) string {
	runes := []rune(phone_number)
	for i := 0; i < len(runes)-4; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

func Test_solution(t *testing.T) {
	type args struct {
		phone_number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{phone_number: "01033334444"}, want: "*******4444"},
		{name: "", args: args{phone_number: "027778888"}, want: "*****8888"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.phone_number); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
