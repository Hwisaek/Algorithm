package _7

func removeElement(nums []int, val int) (result int) {
	for i := len(nums) - 1; i >= 0; i-- {
		e := nums[i]
		if e == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
