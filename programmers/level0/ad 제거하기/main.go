package main

import "regexp"

func solution(strArr []string) (result []string) {
	regex := regexp.MustCompile("ad")
	for _, e := range strArr {
		if regex.FindString(e) == "" {
			result = append(result, e)
		}
	}
	return
}

func main() {
	solution([]string{"and", "notad", "abcd"})
}
