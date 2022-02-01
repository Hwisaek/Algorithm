package 완주하지_못한_선수_test

import "testing"

func solution(participant, completion []string) (result string) {
	pMap := make(map[string]int)
	for _, s := range participant {
		pMap[s]++
	}
	for _, s := range completion {
		pMap[s]--
	}

	for key, value := range pMap {
		if value > 0 {
			result = key
		}
	}
	return
}


func TestSolution(t *testing.T) {
	testCase1 := solution([]string{"leo", "kiki", "eden"}, []string{"eden", "kiki"})
	if testCase1 != "leo" {
		t.Error("testCase1 Error")
	}
	testCase2 := solution([]string{"marina", "josipa", "nikola", "vinko", "filipa"}, []string{"josipa", "filipa", "marina", "nikola"})
	if testCase2 != "vinko" {
		t.Error("testCase1 Error")
	}
	testCase3 := solution([]string{"mislav", "stanko", "mislav", "ana"}, []string{"stanko", "ana", "mislav"})
	if testCase3 != "mislav" {
		t.Error("testCase1 Error")
	}

}
