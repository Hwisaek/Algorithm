package solution

import (
	"strconv"
	"strings"
)

// E: 입장시간, L: 퇴장시간
// 조건: 같은 날짜. E < L

// 비용
// 입장료: 2
// 첫 1시간: 3
// 이후 시간당 4
func Solution(E string, L string) (answer int) {
	answer += 2 + 3

	eHour, eMinute := strings.Split(E, ":")[0], strings.Split(E, ":")[1]
	lHour, lMinute := strings.Split(L, ":")[0], strings.Split(L, ":")[1]

	eH, _ := strconv.Atoi(eHour)
	eM, _ := strconv.Atoi(eMinute)
	lH, _ := strconv.Atoi(lHour)
	lM, _ := strconv.Atoi(lMinute)

	hourDiff := lH - eH
	minDiff := lM - eM + hourDiff*60

	if minDiff > 60 {
		answer += 4 * ((minDiff+60-1)/60 - 1)
	}

	return
}
