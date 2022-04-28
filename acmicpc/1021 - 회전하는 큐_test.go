package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	strArr := strings.Split(scan(rd), " ")
	n, _ := strconv.Atoi(strArr[0])
	m, _ := strconv.Atoi(strArr[1])

	target := make([]int, 0)
	for _, s := range strings.Split(scan(rd), " ") {
		atoi, _ := strconv.Atoi(s)
		target = append(target, atoi)
	}

	_, _ = wr.WriteString(strconv.Itoa(solution1021(n, m, target)))
	_ = wr.Flush()
}

func solution1021(n, m int, target []int) (count int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	size := len(target)
	for size > 0 {
		idx := getIndexOf(&arr, &target)

		sizeArr := len(arr)
		if idx < int(math.Ceil(float64(sizeArr)/2)) {
			arr = append(arr[idx+1:], arr[:idx]...)
			count += idx
		} else {
			arr = append(arr[idx+1:], arr[:idx]...)
			count += sizeArr - idx
		}

		size = len(target)
	}
	return
}

func getIndexOf(arr *[]int, target *[]int) int {
	for i, i2 := range *arr {
		if i2 == (*target)[0] {
			*target = (*target)[1:]
			return i
		}
	}
	return -1
}

func scan(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark1021(b *testing.B) {
	target := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		target[i] = i + 1
	}
	for i := 1; i < b.N; i++ {
		solution1021(1000, 1000, target)
	}
}

func Test_solution1021(t *testing.T) {
	type args struct {
		n      int
		m      int
		target []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "",
			args: args{
				n:      10,
				m:      3,
				target: []int{1, 2, 3},
			},
			wantResult: 0,
		},
		{
			name: "",
			args: args{
				n:      10,
				m:      3,
				target: []int{2, 9, 5},
			},
			wantResult: 8,
		},
		{
			name: "",
			args: args{
				n:      32,
				m:      6,
				target: []int{27, 16, 30, 11, 6, 23},
			},
			wantResult: 59,
		},
		{
			name: "",
			args: args{
				n:      10,
				m:      10,
				target: []int{1, 6, 3, 2, 7, 9, 8, 4, 10, 5},
			},
			wantResult: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1021(tt.args.n, tt.args.m, tt.args.target); gotResult != tt.wantResult {
				t.Errorf("solution1021() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
