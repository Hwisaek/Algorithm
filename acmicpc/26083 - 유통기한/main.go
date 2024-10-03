package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func scan() string {
	s.Scan()
	return s.Text()
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	Y, _ := strconv.Atoi(scan())
	M, _ := strconv.Atoi(scan())
	D, _ := strconv.Atoi(scan())

	N, _ := strconv.Atoi(scan())
	now, _ := time.Parse("20060102", fmt.Sprintf("20%02d%02d%02d", Y, M, D))

	for i := 0; i < N; i++ {
		A, _ := strconv.Atoi(scan())
		B, _ := strconv.Atoi(scan())
		C, _ := strconv.Atoi(scan())

		L := [][3]int{
			{A, B, C},
			{A, C, B},
			{B, A, C},
			{B, C, A},
			{C, A, B},
			{C, B, A},
		}

		result := func() (r string) {
			tList := make([]time.Time, 0, 6)
			for _, l := range L {
				t, err := time.Parse("20060102", fmt.Sprintf("20%02d%02d%02d", l[0], l[1], l[2]))
				if err == nil {
					tList = append(tList, t)
				}
			}

			if len(tList) == 0 {
				return "invalid"
			}

			for _, t := range tList {
				if t.Before(now) {
					return "unsafe"
				}
			}

			return "safe"
		}()

		fmt.Fprintln(w, result)
	}
}

/*
22 11 26
1
99 99 99
invalid

22 11 26
1
22 12 31
safe


*/
