package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
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

	T, _ := strconv.Atoi(scan()) // 테스트 케이스 수 읽기
	for i := 0; i < T; i++ {
		n := scan()                  // 소수 형태의 입력 받기
		n = n[2:]                    // "0." 부분 제거
		idx := strings.Index(n, "(") // 순환 소수 시작 인덱스 찾기

		a, b := 0, 0
		a2, b2 := 1, 1

		switch {
		case idx < 0: // 순환 소수 없음
			ai, _ := strconv.Atoi(n)
			a = ai
			a2 = int(math.Pow10(len(n))) // 소수점 이하 자리수에 해당하는 분모
		case idx == 0: // 순환 소수만 있음
			bi, _ := strconv.Atoi(n[1 : len(n)-1]) // 순환 소수 값
			b = bi
			b2 = int(math.Pow10(len(n[1 : len(n)-1]))) // 순환 소수 길이만큼의 분모
			b2 = b2 - 1
		case idx > 0: // 유한 소수 + 순환 소수
			ai, _ := strconv.Atoi(n[:idx]) // 유한 소수 값
			a = ai
			bi, _ := strconv.Atoi(n[idx+1 : len(n)-1]) // 순환 소수 값
			b = bi
			a2 = int(math.Pow10(len(n[:idx]))) // 전체 분모
			b2 = int(math.Pow10(len(n[idx+1 : len(n)-1])))
			b2 = b2 - 1
			b2 = b2 * int(math.Pow10(idx))
		}

		// 일반분수 + 순환소수분수
		bigA1 := (&big.Int{}).Mul(big.NewInt(int64(a)), big.NewInt(int64(b2)))
		bigA2 := (&big.Int{}).Mul(big.NewInt(int64(b)), big.NewInt(int64(a2)))
		x := bigA1.Add(bigA1, bigA2)
		y := (&big.Int{}).Mul(big.NewInt(int64(a2)), big.NewInt(int64(b2)))

		// 기약분수로 만들기
		gcd := gcd(new(big.Int).Set(x), new(big.Int).Set(y))
		x.Div(x, gcd)

		y.Div(y, gcd)
		fmt.Fprintf(w, "%s/%s\n", x, y) // 결과 출력
	}
}

// GCD 계산 함수
func gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return gcd(b, a.Mod(a, b))
}
