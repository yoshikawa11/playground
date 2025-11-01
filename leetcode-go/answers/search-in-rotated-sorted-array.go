package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid+1] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}
