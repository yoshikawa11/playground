package main

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// 1. 後ろから見て、nums[i] < nums[i+1] となる最後の i を探す
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 2. nums[i] より大きい要素のうち、最も右側のものを探す
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 3. スワップ
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 4. i+1 以降を反転（次の辞書順にする）
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
