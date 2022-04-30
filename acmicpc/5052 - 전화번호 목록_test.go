package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	t, _ := strconv.Atoi(scan5052(rd))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(scan5052(rd))
		numberList := make([]string, 0, n)
		for i := 0; i < n; i++ {
			numberList = append(numberList, scan5052(rd))
		}
		_, _ = wr.WriteString(solution5052(numberList))
	}

	_ = wr.Flush()
}

func solution5052(numberList []string) (result string) {
	sort.Slice(numberList, func(i, j int) bool {
		return numberList[i] < numberList[j]
	})
	lenNumberList := len(numberList)
	for i := 0; i < lenNumberList-1; i++ {
		if strings.HasPrefix(numberList[i+1], numberList[i]) {
			return "NO"
		}
	}
	return "YES"
}

func scan5052(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark5052(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution5052([]string{"113", "12340", "123440", "12345", "98346"})
	}
}

func Test_solution5052(t *testing.T) {
	type args struct {
		numberList []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{numberList: []string{"911", "97625999", "91125426"}}, wantResult: "NO"},
		{name: "", args: args{numberList: []string{"113", "12340", "123440", "12345", "98346"}}, wantResult: "YES"},
		{name: "", args: args{numberList: []string{"123", "12340", "123440", "12345", "98346"}}, wantResult: "NO"},
		{name: "", args: args{numberList: []string{"123", "123440"}}, wantResult: "NO"},
		{name: "", args: args{numberList: []string{"123440", "123"}}, wantResult: "NO"},
		{name: "", args: args{numberList: []string{"113"}}, wantResult: "YES"},
		{name: "", args: args{numberList: []string{"1", "3", "91125426", "911", "97625999"}}, wantResult: "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution5052(tt.args.numberList); gotResult != tt.wantResult {
				t.Errorf("solution5052() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
