package main

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	strArr := strings.Split(scan1764(rd), " ")
	n, _ := strconv.Atoi(strArr[0])
	m, _ := strconv.Atoi(strArr[1])

	nArr := make([]string, 0, n)
	for i := 0; i < n; i++ {
		nArr = append(nArr, scan1764(rd))
	}

	mArr := make([]string, 0, m)
	for i := 0; i < m; i++ {
		mArr = append(mArr, scan1764(rd))
	}

	_, _ = wr.WriteString(strings.Join(solution1764(nArr, mArr), "\n"))
	_ = wr.Flush()
}

func solution1764(nArr, mArr []string) (result []string) {
	result = make([]string, 0, len(nArr))
	count := make(map[string]int)
	for _, s := range nArr {
		count[s]++
	}
	for _, s := range mArr {
		count[s]++
		if count[s] == 2 {
			result = append(result, s)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	result = append([]string{strconv.Itoa(len(result))}, result...)
	return
}

func scan1764(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Test_solution1764(t *testing.T) {
	type args struct {
		nArr []string
		mArr []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "",
			args:       args{nArr: []string{"ohhenrie", "charlie", "baesangwook"}, mArr: []string{"obama", "baesangwook", "ohhenrie", "clinton"}},
			wantResult: []string{"2", "baesangwook", "ohhenrie"},
		},
		{
			name:       "",
			args:       args{nArr: []string{"ohhenrie", "charlie", "baesangwook"}, mArr: []string{"obama", "ohhenrie", "baesangwook", "clinton"}},
			wantResult: []string{"2", "baesangwook", "ohhenrie"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1764(tt.args.nArr, tt.args.mArr); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("solution1764() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
