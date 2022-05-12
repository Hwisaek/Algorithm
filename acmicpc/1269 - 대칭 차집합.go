package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var a, b int
	_, _ = fmt.Fscanln(rd, &a, &b)
	arr1 := strings.Split(scan1269(rd), " ")
	A := make(map[string]struct{})
	for i := 0; i < a; i++ {
		A[arr1[i]] = struct{}{}
	}

	arr2 := strings.Split(scan1269(rd), " ")
	B := make(map[string]struct{})
	for i := 0; i < b; i++ {
		B[arr2[i]] = struct{}{}
	}

	if len(A) > len(B) {
		A, B = B, A
	}
	intersection := 0
	for s := range A {
		_, exist := B[s]
		if exist {
			intersection++
		}
	}

	_, _ = wr.WriteString(strconv.Itoa(len(A) + len(B) - intersection*2))
	_ = wr.Flush()
}

func scan1269(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n')
	str = strings.TrimSpace(str) // 마지막 줄바꿈 문자가 포함되므로 제거
	return str
}
