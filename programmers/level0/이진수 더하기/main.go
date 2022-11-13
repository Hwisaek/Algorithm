package main

import "strconv"

func solution(bin1 string, bin2 string) string {
	parseInt1, _ := strconv.ParseInt(bin1, 2, 64)
	parseInt2, _ := strconv.ParseInt(bin2, 2, 64)
	return strconv.FormatInt(parseInt1+parseInt2, 2)
}
