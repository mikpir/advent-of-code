package d2

func isSliceOkBrute(nums []int) bool {
	if check(nums) {
		return true
	}

	nums2 := make([]int, len(nums)-1)
	for i := range nums {
		index := 0
		for j, n := range nums {
			if j == len(nums) {
				continue
			}
			if i == j {
				continue
			}
			nums2[index] = n
			index++
		}
		if check(nums2) {
			return true
		}
	}
	return false
}

func check(nums []int) bool {
	prev2 := 0
	prev := 0
	for i, curr := range nums {
		if i == 0 {
			prev = curr
			continue
		}

		if !isChangeOk(prev2, prev, curr) {
			return false
		}

		prev2 = prev
		prev = curr
	}
	return true
}
