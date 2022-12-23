import (
	"math/big"
)

func solution(balls int, share int) int {
	return int(new(big.Int).Div(factorial(balls), new(big.Int).Mul(factorial(balls-share), factorial(share))).Int64())
}

func factorial(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(0).Mul(big.NewInt(int64(n)), factorial(n-1))
}
