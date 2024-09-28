package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		cnt := 1

		i := big.NewInt(1)

		for i.Cmp(big.NewInt(int64(n))) < 0 {
			i = i.Mul(i, big.NewInt(int64(10)))
			i = i.Add(i, big.NewInt(1))
			cnt++
		}

		for i.Mod(i, big.NewInt(int64(n))).Cmp(big.NewInt(0)) != 0 {
			cnt++

			i = i.Mul(i, big.NewInt(int64(10)))
			i = i.Add(i, big.NewInt(1))
		}
		fmt.Fprintln(w, cnt)
	}
}
