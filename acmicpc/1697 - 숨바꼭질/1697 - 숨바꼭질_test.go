package _697___숨바꼭질

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanf(rd, "%d %d", &n, &k)

	_, _ = wr.WriteString(fmt.Sprintf("%d\n", solution1697(n, k)))
	_ = wr.Flush()
}

func solution1697(n, k int) (result int) {

	return
}

func bfs(n, k int) (result int) {
	queue := make([]int, 0)
	return
}

func Benchmark_solution5525(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1697(2, 13)
	}
}

func Test_solution5525(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "",
			args: args{
				n: 5,
				k: 17,
			},
			wantResult: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1697(tt.args.n, tt.args.k); gotResult != tt.wantResult {
				t.Errorf("solution1697() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
