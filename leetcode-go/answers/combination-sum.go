package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var comb []int

	// 候補をソートしておけば、残りターゲットが小さいときに打ち切れる
	sort.Ints(candidates)

	var dfs func(startIdx, remaining int)
	dfs = func(startIdx, remaining int) {
		if remaining == 0 {
			// 組み合わせが完成したのでコピーして結果に追加
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			res = append(res, combCopy)
			return
		}
		for i := startIdx; i < len(candidates); i++ {
			v := candidates[i]
			if v > remaining {
				// これ以降の候補も大きいので打ち切り
				break
			}
			// 候補 v を使ってみる
			comb = append(comb, v)
			// 再帰：同じ i をスタートにして v を再度使えるようにする
			dfs(i, remaining-v)
			// バックトラック
			comb = comb[:len(comb)-1]
		}
	}

	dfs(0, target)
	return res
}
