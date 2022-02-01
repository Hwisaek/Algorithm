package 전화번호_목록_test

import (
	"strconv"
	"testing"
)

func solution(phoneBook []string) (result bool) {
	for _, num := range phoneBook {
		n := ""
		for _, i := range num[:len(num)-1] {
			n += strconv.Itoa(int(i) - '0')
			if val, ok := dict["foo"]; ok {
				//do something here
			}
		}
	}
	return
}

func TestSolution(t *testing.T) {
	testCase1 := solution([]string{"119", "97674223", "1195524421"})
	if testCase1 != false {
		t.Error("testCase1 Error")
	}
	testCase2 := solution([]string{"123", "456", "789"})
	if testCase2 != true {
		t.Error("testCase2 Error")
	}
	testCase3 := solution([]string{"12", "123", "1235", "567", "88"})
	if testCase3 != false {
		t.Error("testCase3 Error")
	}
}
