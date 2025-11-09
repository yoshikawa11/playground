package main

func jump(nums []int) int {
	minSteps := 100000
	for i := range nums {
		steps := 0
		position := i
		for position != len(nums)-1 {
			steps++
			if position+nums[position] <= len(nums)-1 {
				position += nums[position]
			} else {
				break
			}
		}

		if position == len(nums)-1 {
			minSteps = min(steps, minSteps)
		}

	}
	return minSteps
}
