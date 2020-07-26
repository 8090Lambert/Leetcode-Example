package _1_search

func Search (nums []int, target int) bool {
	count := len(nums)
	left := 0
	right := count - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[right] && nums[mid] == nums[left] {
			left++
			right--
		} else if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
