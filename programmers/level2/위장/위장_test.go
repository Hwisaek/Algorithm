package 다리를_지나는_트럭

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(clothes [][]string) (result int) {
	result = 1
	set := map[string][]string{}

	for _, clothe := range clothes {
		clotheName := clothe[0]
		clotheType := clothe[1]
		set[clotheType] = append(set[clotheType], clotheName)
	}

	for _, strings := range set {
		result *= len(strings) + 1
	}

	result -= 1
	return
}

func TestSolution(t *testing.T) {
	clothes := [][][]string{{{"yellowhat", "headgear"}, {"bluesunglasses", "eyewear"}, {"green_turban", "headgear"}}, {{"crowmask", "face"}, {"bluesunglasses", "face"}, {"smoky_makeup", "face"}}}
	answer := []int{5, 3}

	for i := range answer {
		result := solution(clothes[i])
		if !reflect.DeepEqual(result, answer[i]) {
			t.Error(fmt.Sprintf("Testcase %v Failed. result: %v, answer: %v", i+1, result, answer[i]))
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	truckWeights := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		truckWeights = append(truckWeights, 10000)
	}
	for i := 0; i < b.N; i++ {
		solution([][]string{{"yellowhat", "headgear"}, {"bluesunglasses", "eyewear"}, {"green_turban", "headgear"}})
	}
}
