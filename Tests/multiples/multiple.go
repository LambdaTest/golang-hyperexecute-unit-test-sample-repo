package multiples

func isMultipleOfFive(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	for _, n := range nums {
		if n%5 != 0 {
			return false
		}
	}

	return true
}