package solution

import (
	"fmt"
	"testing"
)

func Solution(A int, B int, C int) int {
	countA := conformCount(A)
	countB := conformCount(B)
	countC := conformCount(C)

	countAB := conformCount(A | B)
	countAC := conformCount(A | C)
	countBC := conformCount(B | C)

	countABC := conformCount(A | B | C)

	return countA + countB + countC - countAB - countAC - countBC + countABC
}

func conformCount(num int) int {
	zeroCount := 0
	mask := 1
	for i := 0; i < 30; i++ {
		if (num & mask) == 0 {
			zeroCount++
		}
		mask <<= 1
	}

	return 1 << zeroCount
}

func TestSolution(t *testing.T) {
	tests := []struct {
		A, B, C int
		want    int
	}{
		{1073741727, 1073741631, 1073741679, 8},
		{16244239, 13032961, 819399173, 0},
		{0, 0, 0, 1},
		{1073741823, 1073741823, 1073741823, 1},
		{1, 2, 4, 7},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("A=%d, B=%d, C=%d", tt.A, tt.B, tt.C), func(t *testing.T) {
			result := Solution(tt.A, tt.B, tt.C)
			if result != tt.want {
				t.Errorf("expected %d, got %d", tt.want, result)
			}
		})
	}
}
