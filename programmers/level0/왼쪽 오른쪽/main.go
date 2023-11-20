package main

func solution(str_list []string) []string {
	for i, s := range str_list {
		switch s {
		case "l":
			return str_list[:i]
		case "r":
			return str_list[i+1:]
		}
	}
	return []string{}
}
