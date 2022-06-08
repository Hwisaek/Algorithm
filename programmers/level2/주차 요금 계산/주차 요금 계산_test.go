package main

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func solution(fees []int, records []string) (result []int) {
	baseTime, baseRate, unitTime, unitFee := fees[0], fees[1], fees[2], fees[3]

	inOut := make(map[string][]*int)

	for _, record := range records {
		split := strings.Split(record, " ")
		hhmm, car, flag := split[0], split[1], split[2]
		split2 := strings.Split(hhmm, ":")
		h, _ := strconv.Atoi(split2[0])
		m, _ := strconv.Atoi(split2[1])
		time := h*60 + m
		if _, exist := inOut[car]; !exist {
			inOut[car] = make([]*int, 3)
			zero := 0
			inOut[car][2] = &zero
		}
		if flag == "IN" {
			inOut[car][0] = &time
		} else {
			inOut[car][1] = &time
			*inOut[car][2] += *inOut[car][1] - *inOut[car][0]
			inOut[car][0] = nil
			inOut[car][1] = nil
		}
	}

	numFeeArr := make([][]int, 0, len(inOut))
	for car, arr := range inOut {
		carNum, _ := strconv.Atoi(car)

		if arr[0] != nil {
			*arr[2] += (23*60 + 59) - *arr[0]
		}
		//주차요금 = 기본요금 + ⌈(누적 주차 시간 - 기본 시간) / 단위 시간⌉ x 단위 요금
		diff := *arr[2] - baseTime
		if diff < 0 {
			diff = 0
		}
		fee := baseRate + int(math.Ceil(float64(diff)/float64(unitTime)))*unitFee

		numFeeArr = append(numFeeArr, []int{carNum, fee})
	}

	sort.Slice(numFeeArr, func(i, j int) bool {
		return numFeeArr[i][0] < numFeeArr[j][0]
	})
	for _, arr := range numFeeArr {
		result = append(result, arr[1])
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		fees    []int
		records []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				fees:    []int{180, 5000, 10, 600},
				records: []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"},
			}, want: []int{14600, 34400, 5000}},
		{
			name: "",
			args: args{
				fees:    []int{120, 0, 60, 591},
				records: []string{"16:00 3961 IN", "16:00 0202 IN", "18:00 3961 OUT", "18:00 0202 OUT", "23:58 3961 IN"},
			}, want: []int{0, 591}},
		{
			name: "",
			args: args{
				fees:    []int{1, 461, 1, 10},
				records: []string{"00:00 1234 IN"},
			}, want: []int{14841}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.fees, tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]int{180, 5000, 10, 600}, []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"})
	}
}
