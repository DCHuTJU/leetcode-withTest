package findMagicIndex

func FindMagicIndex(nums []int) int {
	return findSubMagicIndex(nums, 0, len(nums) - 1)
}

func findSubMagicIndex(nums []int, low, high int) int {
	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] == mid {
			ret := findSubMagicIndex(nums, low, mid-1)
			if ret < mid && ret != -1 {
				return ret
			} else {
				return mid
			}
		} else {
			ret := findSubMagicIndex(nums, low, mid - 1)
			if ret != -1 {
				return ret
			} else {
				return findSubMagicIndex(nums, mid + 1, high)
			}
		}
	}
	return -1
}

func FindMagicIndexII(nums []int) int {
	return getAnswer(nums, 0, len(nums) - 1)
}

func getAnswer(nums []int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (right - left) / 2 + left
	leftAnswer := getAnswer(nums, left, mid - 1)
	if leftAnswer != -1 {
		return leftAnswer
	} else if nums[mid] == mid {
		return mid
	}
	return getAnswer(nums, mid + 1, right)
}