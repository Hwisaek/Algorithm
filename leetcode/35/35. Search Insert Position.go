package _5

func searchInsert(nums []int, target int) (result int) {
	startIdx := 0
	endIdx := len(nums)

	for {
		medianIdx := (startIdx + endIdx) / 2
		median := nums[medianIdx]

		if median == target {
			result = medianIdx
			break
		} else if median > target {
			if endIdx-startIdx == 1 {
				result = medianIdx - 1
				break
			}
			endIdx = medianIdx
		} else {
			if endIdx-startIdx == 1 {
				result = medianIdx + 1
				break
			}
			startIdx = medianIdx
		}
	}
	if result < 0 {
		result = 0
	}
	return
}
