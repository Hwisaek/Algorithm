package 전화번호_목록_test

import (
	"testing"
)

func solution(phoneBook []string) bool {
	dict := make(map[string]string)
	for _, s := range phoneBook {
		dict[s] = s
	}

	for _, num := range phoneBook {
		n := ""
		for i := range num[:len(num)-1] {
			n += num[i : i+1]
			if _, ok := dict[n]; ok {
				return false
			}
		}
	}
	return true
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
