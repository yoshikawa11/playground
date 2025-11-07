package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int

	sort.Ints(candidates)

	var dfs func(startIdx, remaining int)
	dfs = func(startIdx, remaining int) {
		if remaining == 0 {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			res = append(res, combCopy)
			return
		}

		for i := startIdx; i < len(candidates); i++ {
			if i > startIdx && candidates[i] == candidates[i-1] {
				continue
			}
			v := candidates[i]
			if v > remaining {
				break
			}

			comb = append(comb, v)
			dfs(i+1, remaining-v)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(0, target)
	return res
}
