package main

import "strings"

func solution(survey []string, choices []int) (result string) {
	a := []int{0, 0, 0, 0}
	for i, s := range survey {
		switch {
		case strings.ContainsAny(s, "RT"):
			if s[0] == 'R' {
				a[0] += choices[i] - 4
			} else {
				a[0] -= choices[i] - 4
			}
		case strings.ContainsAny(s, "CF"):
			if s[0] == 'C' {
				a[1] += choices[i] - 4
			} else {
				a[1] -= choices[i] - 4
			}
		case strings.ContainsAny(s, "JM"):
			if s[0] == 'J' {
				a[2] += choices[i] - 4
			} else {
				a[2] -= choices[i] - 4
			}
		case strings.ContainsAny(s, "AN"):
			if s[0] == 'A' {
				a[3] += choices[i] - 4
			} else {
				a[3] -= choices[i] - 4
			}
		}
	}

	if a[0] <= 0 {
		result += "R"
	} else {
		result += "T"
	}
	if a[1] <= 0 {
		result += "C"
	} else {
		result += "F"
	}
	if a[2] <= 0 {
		result += "J"
	} else {
		result += "M"
	}
	if a[3] <= 0 {
		result += "A"
	} else {
		result += "N"
	}
	return
}

func main() {
	solution([]string{"AN", "CF", "MJ", "RT", "NA"}, []int{5, 3, 2, 7, 5})
	solution([]string{"TR", "RT", "TR"}, []int{7, 1, 3})
}
