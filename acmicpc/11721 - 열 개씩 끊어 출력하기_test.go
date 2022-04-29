package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	word := scan11721(rd)

	_, _ = wr.WriteString(solution11721(word))
	_ = wr.Flush()
}

func solution11721(word string) (result string) {
	strArr := make([]string, 0, len(word)/10+1)

	for i := 0; i < len(word)/10; i++ {
		strArr = append(strArr, word[i*10:(i+1)*10])
	}
	strArr = append(strArr, word[len(word)/10*10:])

	result = strings.Join(strArr, "\n")
	return
}

func scan11721(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark11721(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution11721("OneTwoThreeFourFiveSixSevenEightNineTenOneTwoThreeFourFiveSixSevenEightNineTenOneTwoThreeFourFiveSix")
	}
}

func Test_solution11721(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{word: "BaekjoonOnlineJudge"}, wantResult: "BaekjoonOn\nlineJudge"},
		{name: "", args: args{word: "OneTwoThreeFourFiveSixSevenEightNineTen"}, wantResult: "OneTwoThre\neFourFiveS\nixSevenEig\nhtNineTen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution11721(tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution11721() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
