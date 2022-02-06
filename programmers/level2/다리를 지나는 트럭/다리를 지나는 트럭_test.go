package 다리를_지나는_트럭

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(bridgeLength int, weight int, truckWeights []int) int {
	return 0
}

func TestSolution(t *testing.T) {
	bridgeLength := []int{2, 100, 100}
	weight := []int{10, 100, 100}
	truckWeights := [][]int{{7, 4, 5, 6}, {10}, {10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}
	answer := []int{1, 5}

	for i := range answer {
		result := solution(bridgeLength[i], weight[i], truckWeights[i])
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
		solution(10000, 10000, truckWeights)
	}
}
