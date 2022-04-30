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

	n, _ := strconv.Atoi(scan1302(rd))

	books := make([]string, 0, n)
	for i := 0; i < n; i++ {
		books = append(books, scan1302(rd))
	}

	_, _ = wr.WriteString(solution1302(books))
	_ = wr.Flush()
}

func solution1302(books []string) string {
	count := make(map[string]int)
	for _, name := range books {
		count[name]++
	}

	maxCount := -1
	maxBooks := make([]string, 0, len(count))
	for name, num := range count {
		if num > maxCount {
			maxBooks = append(maxBooks[:0], name)
			maxCount = num
		} else if num == maxCount {
			maxBooks = append(maxBooks, name)
		}
	}

	sort.Slice(maxBooks, func(i, j int) bool {
		return maxBooks[i] < maxBooks[j]
	})
	return maxBooks[0]
}

func scan1302(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark1302(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1302([]string{"icecream", "peanuts", "peanuts", "chocolate", "candy", "chocolate", "icecream", "apple"})
	}
}
func Test_solution1302(t *testing.T) {
	type args struct {
		books []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{books: []string{"top", "top", "top", "top", "kimtop"}}, wantResult: "top"},
		{name: "", args: args{books: []string{"table", "chair", "table", "table", "lamp", "door", "lamp", "table", "chair"}}, wantResult: "table"},
		{name: "", args: args{books: []string{"icecream", "peanuts", "peanuts", "chocolate", "candy", "chocolate", "icecream", "apple"}}, wantResult: "chocolate"},
		{name: "", args: args{books: []string{"soul"}}, wantResult: "soul"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1302(tt.args.books); gotResult != tt.wantResult {
				t.Errorf("solution1302() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
