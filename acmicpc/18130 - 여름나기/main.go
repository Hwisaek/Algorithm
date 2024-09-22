package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
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

	N := new(big.Int)
	N.SetString(scan(), 10)
	Q := new(big.Int)
	Q.SetString(scan(), 10)

	idx := 1
	minCost := new(big.Int)
	minCost.SetString("1000000000000000000000000000000", 10) // 매우 큰 수로 초기화

	two := big.NewInt(2)

	for i := 1; i <= int(N.Int64()); i++ {
		P := new(big.Int)
		P.SetString(scan(), 10)
		K := new(big.Int)
		K.SetString(scan(), 10)
		C := new(big.Int)
		C.SetString(scan(), 10)

		floorQK := new(big.Int).Div(Q, K)

		sum := new(big.Int).Add(floorQK, big.NewInt(1))
		sum.Mul(sum, floorQK)
		sum.Div(sum, two)

		totalCost := new(big.Int).Mul(C, sum)
		totalCost.Add(totalCost, P)

		if totalCost.Cmp(minCost) < 0 {
			minCost.Set(totalCost)
			idx = i
		}
	}

	fmt.Fprintf(w, "%d %s\n", idx, minCost.String())
}
