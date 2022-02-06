package 다리를_지나는_트럭

import (
	"fmt"
	"reflect"
	"testing"
)

func solution(bridgeLength int, weight int, truckWeights []int) (time int) {
	bridge := make([]int, bridgeLength)
	curWeight := 0
	for len(truckWeights) > 0 || curWeight > 0 {
		curWeight -= bridge[0]
		bridge = bridge[1:]
		if len(truckWeights) > 0 {
			if curWeight+truckWeights[0] <= weight {
				truck := truckWeights[0]
				truckWeights = truckWeights[1:]
				bridge = append(bridge, truck)
				curWeight += truck
			} else {
				bridge = append(bridge, 0)
			}
		}
		time++
	}
	return
}

func TestSolution(t *testing.T) {
	bridgeLength := []int{2, 100, 100}
	weight := []int{10, 100, 100}
	truckWeights := [][]int{{7, 4, 5, 6}, {10}, {10, 10, 10, 10, 10, 10, 10, 10, 10, 10}}
	answer := []int{8, 101, 110}

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
