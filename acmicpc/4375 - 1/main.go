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

		a, b := big.NewInt(0), big.NewInt(0)
		a2, b2 := big.NewInt(1), big.NewInt(1)

		switch {
		case idx < 0: // 순환 소수 없음
			ai, _ := strconv.Atoi(n)
			a = big.NewInt(int64(ai))
			a2 = big.NewInt(int64(math.Pow10(len(n)))) // 소수점 이하 자리수에 해당하는 분모
		case idx == 0: // 순환 소수만 있음
			bi, _ := strconv.Atoi(n[1 : len(n)-1]) // 순환 소수 값
			b = big.NewInt(int64(bi))
			b2 = big.NewInt(int64(math.Pow10(len(n[1 : len(n)-1])))) // 순환 소수 길이만큼의 분모
			b2 = b2.Sub(b2, big.NewInt(1))
		case idx > 0: // 유한 소수 + 순환 소수
			ai, _ := strconv.Atoi(n[:idx]) // 유한 소수 값
			a = big.NewInt(int64(ai))
			bi, _ := strconv.Atoi(n[idx+1 : len(n)-1]) // 순환 소수 값
			b = big.NewInt(int64(bi))
			a2 = big.NewInt(int64(math.Pow10(len(n)))) // 전체 분모
			b2 = big.NewInt(int64(math.Pow10(len(n[idx+1 : len(n)-1]))))
			b2 = b2.Sub(b2, big.NewInt(1))
		}

		// 일반분수 + 순환소수분수
		var x, y *big.Int
		x = new(big.Int).Set(a)                  // a를 x로 초기화
		x.Mul(x, b2)                             // x에 b2 곱하기
		x.Add(x, new(big.Int).Set(b).Mul(b, a2)) // x에 b * a2 더하기
		y = new(big.Int).Set(a2)                 // y에 a2 초기화
		y.Mul(y, b2)                             // y에 b2 곱하기

		// 기약분수로 만들기
		gcd := gcd(x, y)

		x.Div(x, gcd)
		y.Div(y, gcd)
		fmt.Fprintf(w, "%s/%s\n", x.Text(10), y.Text(10)) // 결과 출력
	}
}

// GCD 계산 함수
func gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return gcd(b, new(big.Int).Mod(a, b))
}
